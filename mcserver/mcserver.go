package mcserver

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"github.com/initsuj/gomc/protocol"
	"net"
	"strconv"
)

const (
	DefaultPort uint16 = 25565
)

type StatusResponse struct {
	Response string
}

type Status struct {
	Desc struct {
		Text string `json:"text"`
	} `json:"description"`
	Players struct {
		Max    int `json:"max"`
		Online int `json:"online"`
	} `json:"players"`
	Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	} `json:"version"`
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

	addr, err := net.ResolveTCPAddr("tcp", server)
	if err != nil {
		return err
	}

	h := protocol.NewStatusHandshakePkt(s, port)

	fmt.Printf("pkt: %v\n", h)

	pkt, err := h.MarshalBinary()
	if err != nil {
		return err
	}

	var buf bytes.Buffer
	buf.Write(pkt)

	c, err := net.DialTCP("tcp", nil, addr)
	if err != nil {
		return err
	}

	defer c.Close()

	fmt.Printf("sent: %02x\n", buf.Bytes())
	buf.WriteTo(c)

	buf.Reset()

	ping := protocol.NewEmptyRequest()
	pkt, err = ping.MarshalBinary()
	if err != nil {
		return err
	}

	buf.Write(pkt)
	fmt.Printf("sent: %02x\n", buf.Bytes())
	buf.WriteTo(c)

	buf.ReadFrom(c)

	fmt.Printf("received: %x\n", buf.Bytes())

	response := protocol.NewStatusResponse()
	err = response.UnmarshalBinary(&buf)
	if err != nil {
		return err
	}

	fmt.Printf("response: %v\n", response.Id)

	/*
		id, n := binary.Uvarint(buf.Bytes())
		if n < 0 {
			return errors.New("Error getting packet length!")
		}
	*/

	//fmt.Printf("id(%v) ", id)

	/*
		if err := bin.NewDecoder(c).Decode(&response); err != nil {
			return err
		}

		fmt.Printf("received: %v\n", response.Response)

		var status Status
		if err := json.Unmarshal([]byte(response.Response), &status); err != nil {
			return err
		}

		fmt.Printf("received: %v\n", status.Desc.Text)
	*/

	return nil

}

func varint2b(x uint64) []byte {
	b := make([]byte, binary.MaxVarintLen64)
	n := binary.PutUvarint(b, x)
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
