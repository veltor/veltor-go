// Copyright (c) 2017 The Alvalor Authors
//
// This file is part of Alvalor.
//
// Alvalor is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// Alvalor is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with Alvalor.  If not, see <http://www.gnu.org/licenses/>.

package trie

import (
	"errors"
	"hash"

	"golang.org/x/crypto/blake2b"
)

// A list of errors that can be returned by the functions of the trie.
var (
	ErrAlreadyExists = errors.New("key already exists")
	ErrNotFound      = errors.New("key not found")
)

// Trie represents our own implementation of the patricia merkle trie as specified in the Ethereum
// yellow paper, with a few simplications due to the simpler structure of the Alvalor blockchain.
type Trie struct {
	root node
	h    hash.Hash
}

// New creates a new empty trie with no state.
func New() *Trie {
	h, _ := blake2b.New256(nil)
	t := &Trie{h: h}
	return t
}

// Put will insert the given data for the given key. It will fail if there already is data with the given key.
func (t *Trie) Put(key []byte, data []byte) error {
	return t.put(key, data, false)
}

// MustPut will insert the given data for the given key and will overwrite any data that might already be stored under
// the given key.
func (t *Trie) MustPut(key []byte, data []byte) {
	t.put(key, data, true)
}

func (t *Trie) put(key []byte, data []byte, force bool) error {
	cur := &t.root
	path := encode(key)
	for {
		switch n := (*cur).(type) {
		case *fullNode:
			cur = &n.children[path[0]]
			path = path[1:]
		case *shortNode:
			var common []byte
			for i := 0; i < len(n.key); i++ {
				if path[i] != n.key[i] {
					break
				}
				common = append(common, path[i])
			}
			if len(common) == len(n.key) {
				cur = &n.child
				path = path[len(common):]
				continue
			}
			path = path[len(common):]
			remain := n.key[len(common):]
			var left node
			if len(remain) == 1 {
				left = n.child
			} else {
				left = &shortNode{key: remain[1:], child: n.child}
			}
			full := &fullNode{}
			full.children[remain[0]] = left
			if len(common) > 0 {
				short := &shortNode{key: common, child: full}
				*cur = short
				cur = &short.child
			} else {
				*cur = full
				var next node = full
				cur = &next
			}
		case valueNode:
			if !force {
				return ErrAlreadyExists
			}
			*cur = nil
		case nil:
			if len(path) > 0 {
				short := &shortNode{key: path}
				*cur = short
				cur = &short.child
				path = nil
				continue
			}
			*cur = valueNode(data)
			return nil
		}
	}
}

// Get will retrieve the hash located at the path provided by the given key. If the path doesn't
// exist or there is no hash at the given location, it returns a nil slice and false.
func (t *Trie) Get(key []byte) ([]byte, error) {
	cur := &t.root
	path := encode(key)
	for {
		switch n := (*cur).(type) {
		case *fullNode:
			cur = &n.children[path[0]]
			path = path[1:]
		case *shortNode:
			var common []byte
			for i := 0; i < len(n.key); i++ {
				if path[i] != n.key[i] {
					break
				}
				common = append(common, path[i])
			}
			if len(common) == len(n.key) {
				cur = &n.child
				path = path[len(common):]
				continue
			}
			return nil, ErrNotFound
		case valueNode:
			return []byte(n), nil
		case nil:
			return nil, ErrNotFound
		}
	}
}

// Del will try to delete the hash located at the path provided by the given key. If no hash is
// found at the given location, it returns false.
func (t *Trie) Del(key []byte) error {
	var visited []*node
	cur := &t.root
	path := encode(key)
Remove:
	for {
		switch n := (*cur).(type) {
		case *fullNode:
			visited = append(visited, cur)
			cur = &n.children[path[0]]
			path = path[1:]
		case *shortNode:
			visited = append(visited, cur)
			var common []byte
			for i := 0; i < len(n.key); i++ {
				if path[i] != n.key[i] {
					break
				}
				common = append(common, path[i])
			}
			if len(common) == len(n.key) {
				cur = &n.child
				path = path[len(common):]
				continue
			}
			return ErrNotFound
		case valueNode:
			*cur = nil
			break Remove
		case nil:
			return ErrNotFound
		}
	}
Compact:
	for len(visited) > 0 {
		cur = visited[len(visited)-1]
		switch n := (*cur).(type) {
		case *shortNode:
			*cur = nil
			visited = visited[:len(visited)-1]
			continue Compact
		case *fullNode:
			var index int
			var child node
			count := 0
			for i, c := range n.children {
				if c != nil {
					index = i
					child = c
					count++
				}
			}
			if count > 1 {
				break Compact
			}
			short := shortNode{
				key:   []byte{byte(index)},
				child: child,
			}
			c, ok := child.(*shortNode)
			if ok {
				short.key = append(short.key, c.key...)
				short.child = c.child
			}
			*cur = &short
			if len(visited) > 1 {
				parent := visited[len(visited)-2]
				p, ok := (*parent).(*shortNode)
				if ok {
					p.key = append(p.key, short.key...)
					p.child = short.child
				}
			}
			break Compact
		}
	}
	return nil
}

// Hash will return the hash that represents the trie in its entirety by returning the hash of the
// root node. Currently, it does not do any caching and recomputes the hash from the leafs up. If
// the root is not initialized, it will return the hash of an empty byte array to uniquely represent
// a trie without state.
func (t *Trie) Hash() []byte {
	return t.nodeHash(t.root)
}

// nodeHash will return the hash of a given node.
func (t *Trie) nodeHash(node node) []byte {
	switch n := node.(type) {
	case *fullNode:
		var hashes [][]byte
		for _, child := range n.children {
			hashes = append(hashes, t.nodeHash(child))
		}
		t.h.Reset()
		for _, hash := range hashes {
			t.h.Write(hash)
		}
		return t.h.Sum(nil)
	case *shortNode:
		hash := t.nodeHash(n.child)
		t.h.Reset()
		t.h.Write(n.key)
		t.h.Write(hash)
		return t.h.Sum(nil)
	case valueNode:
		t.h.Reset()
		t.h.Write([]byte(n))
		return t.h.Sum(nil)
	case nil:
		t.h.Reset()
		return t.h.Sum(nil)
	default:
		panic("invalid node type")
	}
}
