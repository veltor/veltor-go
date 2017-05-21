// Code generated by capnpc-go.

package protocol

// AUTO GENERATED - DO NOT EDIT

import (
	capnp "zombiezen.com/go/capnproto2"
	text "zombiezen.com/go/capnproto2/encoding/text"
	schemas "zombiezen.com/go/capnproto2/schemas"
)

type Pong struct{ capnp.Struct }

// Pong_TypeID is the unique identifier for the type Pong.
const Pong_TypeID = 0xe5b532ff4fa02dff

func NewPong(s *capnp.Segment) (Pong, error) {
	st, err := capnp.NewStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Pong{st}, err
}

func NewRootPong(s *capnp.Segment) (Pong, error) {
	st, err := capnp.NewRootStruct(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0})
	return Pong{st}, err
}

func ReadRootPong(msg *capnp.Message) (Pong, error) {
	root, err := msg.RootPtr()
	return Pong{root.Struct()}, err
}

func (s Pong) String() string {
	str, _ := text.Marshal(0xe5b532ff4fa02dff, s.Struct)
	return str
}

func (s Pong) Nonce() uint32 {
	return s.Struct.Uint32(0)
}

func (s Pong) SetNonce(v uint32) {
	s.Struct.SetUint32(0, v)
}

// Pong_List is a list of Pong.
type Pong_List struct{ capnp.List }

// NewPong creates a new list of Pong.
func NewPong_List(s *capnp.Segment, sz int32) (Pong_List, error) {
	l, err := capnp.NewCompositeList(s, capnp.ObjectSize{DataSize: 8, PointerCount: 0}, sz)
	return Pong_List{l}, err
}

func (s Pong_List) At(i int) Pong { return Pong{s.List.Struct(i)} }

func (s Pong_List) Set(i int, v Pong) error { return s.List.SetStruct(i, v.Struct) }

// Pong_Promise is a wrapper for a Pong promised by a client call.
type Pong_Promise struct{ *capnp.Pipeline }

func (p Pong_Promise) Struct() (Pong, error) {
	s, err := p.Pipeline.Struct()
	return Pong{s}, err
}

const schema_cf51decd769a7cb5 = "x\xda\x12X\xe1\xc0d\xc8Z\xcf\xc2\xc0\x10h\xc2\xca" +
	"\xf6\xff\xbf\xee\x02\xff\xffF[\x9f2\x04\x8a02\xfe" +
	"\xdfZ3\xab\xec\xec\xbd\xc0\xf3\x0c,\xec\x0c\x0c\xc2\xaa" +
	"L\xa7\x84\x0d\x99@,]&{\x86\xd6\xff\x05E\xf9" +
	"%\xf9\xc9\xf99L\xfa\x05\xf9y\xe9z\xc9\x89\x05y" +
	"\x05V\x01@&\x03C\x00#c \x0b3\xd0P\x16" +
	"F\x06\x06A^#\xa0\xe9\x1c\xcc\x8c\x81\"L\x8c\xf2" +
	"y\xf9y\xc9\xa9\x8c\x1c\x0cL@\xcc\x08\x08\x00\x00\xff" +
	"\xff=\xa0\x1fp"

func init() {
	schemas.Register(schema_cf51decd769a7cb5,
		0xe5b532ff4fa02dff)
}
