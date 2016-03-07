package mcauth

import (
	"bytes"
	"crypto/rand"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"github.com/initsuj/gomc/mcauth/mcrequest"
)

var (
	header = &http.Header{}
)

func init() {
	header.Set("User-Agent", "GoMC/1.0")
	header.Set("Content-Type", "application/json")
}

type authError struct {
	Error   string `json:"error"`
	Message string `json:"errorMessage"`
	Cause   string `json:"cause, omitempty"`
}

func Login(l mcrequest.Login, acct *Account) error {
	body, err := json.MarshalIndent(l, "", "    ")
	if err != nil {
		return err
	}
	req, err := http.NewRequest("POST", "https://authserver.mojang.com/authenticate", bytes.NewBuffer(body))
	if err != nil {
		return err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	fmt.Println(string(body[:]))

	if resp.StatusCode == http.StatusOK {
		if json.Unmarshal(body, acct) != nil {
			return err
		}

		return nil
	} else {
		var authErr authError
		if json.Unmarshal(body, &authErr) != nil {
			return err
		}

		return errors.New(authErr.Message)
	}

}

// newUUID generates a random UUID
func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}

	return string(uuid[:]), nil
}
