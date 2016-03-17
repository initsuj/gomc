package protocol

import (
	"bytes"
	"encoding/binary"
	"io"
)

type Handshake struct {
	Version int32
	Addr    string
	Port    uint16
	Next    int32
}

func (h Handshake) MarshalBinary() ([]byte, error) {
	buf := new(bytes.Buffer)

	_, err := buf.Write(varInt2B(h.Version))
	if err != nil {
		return nil, err
	}

	_, err = buf.Write(s2b(h.Addr))
	if err != nil {
		return nil, err
	}

	err = binary.Write(buf, binary.BigEndian, h.Port)
	if err != nil {
		return nil, err
	}

	_, err = buf.Write(varInt2B(h.Next))
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

func (h Handshake) UnmarshalBinary(reader io.ByteReader) (err error) {

	h.Version, err = b2VarInt(reader)
	if err != nil {
		return err
	}

	return nil
}

func NewStatusHandshakePkt(server string, port uint16) *Packet {
	return &Packet{
		Id: 0,
		Payload: Handshake{
			Version: 4,
			Addr:    server,
			Port:    port,
			Next:    1,
		},
	}
}
