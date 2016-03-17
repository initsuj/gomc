package protocol

import (
	"bytes"
	"errors"
	"fmt"
	"io"
)

var ()

func newVarIntError(i int, field string) error {

	if i < 0 {
		return errors.New(fmt.Sprintf("Error in reading VarInt: Buffer overflow in %v!", field))
	}

	return errors.New(fmt.Sprintf("Error in reading VarInt: Buffer underflow in %v!", field))
}

type McPacketizer interface {
	MarshalBinary() (data []byte, err error)
	UnmarshalBinary(reader io.ByteReader) error
}

type Packet struct {
	Id      int32
	Payload McPacketizer
}

func (p *Packet) MarshalBinary() ([]byte, error) {
	var buf bytes.Buffer
	id := varInt2B(p.Id)
	_, err := buf.Write(id)

	if p.Payload != nil {
		payload, err := p.Payload.MarshalBinary()
		if err != nil {
			return buf.Bytes(), err
		}
		_, err = buf.Write(payload)
	}

	l := int32(buf.Len())

	return append(varInt2B(l), buf.Bytes()...), err
}

func (p *Packet) UnmarshalBinary(reader io.ByteReader) error {

	_, err := b2VarInt(reader)
	if err != nil {
		return err
	}

	p.Id, err = b2VarInt(reader)
	if err != nil {
		return err
	}

	if p.Payload != nil {
		p.UnmarshalBinary(reader)
	}

	return nil
}

func NewEmptyRequest() *Packet {
	return &Packet{Id: 0, Payload: nil}
}

func NewStatusResponse() *Packet {
	return &Packet{Payload: &Status{}}
}
