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

package node

type chainManager interface {
	Height() (uint32, error)
	BestHash() ([]byte, error)
	HashByHeight(height uint32) ([]byte, error)
}

type simpleChain struct {
}

func newChain() *simpleChain {
	return &simpleChain{}
}

func (b *simpleChain) Height() (uint32, error) {
	// TODO
	return 0, nil
}

func (b *simpleChain) BestHash() ([]byte, error) {
	// TODO
	return nil, nil
}

func (b *simpleChain) HashByHeight(height uint32) ([]byte, error) {
	// TODO
	return nil, nil
}
