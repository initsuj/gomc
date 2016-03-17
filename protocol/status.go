package protocol

import "io"

type Status struct {
	Desc struct {
		Text string `json:"text"`
	} `json:"description"`
	Players struct {
		Max     int `json:"max"`
		Current int `json:"online"`
		Online  []struct {
			Name string `json:"name"`
			Id   string `json:"id"`
		}
	} `json:"players"`
	Version struct {
		Name     string `json:"name"`
		Protocol int    `json:"protocol"`
	} `json:"version"`
}

func (s Status) MarshalBinary() ([]byte, error) {
	return nil, nil
}

func (s Status) UnmarshalBinary(reader io.ByteReader) error {
	return nil
}
