package protocol

import (
	"encoding/binary"
	"io"
)

type protocol int32

const (
	Protocol18x protocol = 47
	Protocol19x protocol = 107
)

var (
	DefaultProtocol = Protocol19x
)

func varInt2B(v int32) []byte {
	return varLong2B(int64(v))
}

func b2VarInt(b io.ByteReader) (int32, error) {
	v, err := b2VarLong(b)

	return int32(v), err
}

func varLong2B(v int64) []byte {
	buf := make([]byte, 10)
	l := binary.PutUvarint(buf, uint64(v))

	return buf[:l]
}

func b2VarLong(b io.ByteReader) (uint64, error) {
	return binary.ReadUvarint(b)
}

func s2b(s string) []byte {
	b := []byte(s)
	v := varInt2B(int32(len(b)))

	return append(v, []byte(s)...)
}

func b2s(reader io.ByteReader) (string, error) {
	n, err := b2VarInt(reader)
	if err != nil {
		return "", err
	}

	if n == 0 {
		return "", nil
	}

	buf := make([]byte, n)
	for i := 0; i < int(n); i++ {
		buf[i], err = reader.ReadByte()
		if err != nil {
			return "", err
		}
	}

	return string(buf), nil
}
