package mcserver

import (
	"bytes"
	"encoding/binary"
	"fmt"
	bin "github.com/initsuj/binary"
	"io"
	"net"
	"strconv"
)

const (
	DefaultPort uint16 = 25565
)

type Handshake struct {
	Id    int32 `bin:"varint"`
	Proto int32 `bin:"varint"`
	Addr  string
	Port  uint16
	Next  int32 `bin:"varint"`
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

	h := &Handshake{
		Id:    0,
		Proto: 4,
		Addr:  s,
		Port:  port,
		Next:  1,
	}

	pkt, err := bin.Marshal(h)
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	buf.Write(varint2b(uint64(len(pkt))))
	buf.Write(pkt)
	fmt.Printf("sent pkt(%v) %v\n", varint2b(uint64(len(pkt))), h)

	c, err := net.Dial("tcp", server)
	if err != nil {
		return err
	}

	defer c.Close()
	fmt.Printf("listening: %v\n", c.LocalAddr().String())

	fmt.Printf("sent: %02x", buf.Bytes())
	//buf.WriteTo(os.Stdout)
	buf.WriteTo(c)

	buf.Reset()

	ping := varint2b(0)
	buf.Write(varint2b(uint64(len(ping))))
	buf.Write(pkt)

	buf.WriteTo(c)

	var b bytes.Buffer
	io.Copy(&b, c)
	fmt.Println("total size:", b.Len())

	return nil

}

func varint2b(x uint64) []byte {
	fmt.Printf("received %v\n", x)
	b := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(b, x)
	fmt.Printf("created %v : %v\n", n, b[:n])
	return b[:n]
}

func ushort2b(x uint16) (b []byte) {
	b = make([]byte, 2)
	binary.BigEndian.PutUint16(b, x)

	return
}

func str2b(s string) (b []byte) {
	b = make([]byte, len(s))
	b = []byte(s)
	b = append(varint2b(uint64(len(b))), b...)
	return
}
