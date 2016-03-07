package mcserver

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"net"
	"strconv"
)

const (
	DefaultPort uint16 = 25565
)

type packet struct {
	Length int
	Id     uint16
	Data   []byte
}

func GetServerInfo(server string) error {
	s, p, err := net.SplitHostPort(server)
	if err != nil {
		return err
	}

	var port uint16 = DefaultPort
	if p != "" {
		i, err := strconv.ParseUint(p, 10, 16)
		if err != nil {
			return err
		}
		port = uint16(i)
	}

	buf := new(bytes.Buffer)

	//n := binary.PutVarint(buf, 0x00)
	n := binary.PutVarint(buf, 0)
	n += binary.PutVarint(buf, int64(len(s)))
	i, err := buf.WriteString(s)
	if err != nil {
		return err
	}

	n += i
	n += binary.BigEndian.PutUint16(buf, port)
	n += binary.PutVarint(buf, 1)

	fmt.Printf("count: %v - % x", len(buf.Bytes()), buf.Bytes())
	return nil

}
