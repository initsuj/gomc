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
)

var (
	header = &http.Header{}
)

func init() {
	header.Set("User-Agent", "GoMC/1.0")
	header.Set("Content-Type", "application/json")
}

type AuthError struct {
	Type    string `json:"error"`
	Message string `json:"errorMessage"`
	Cause   string `json:"cause, omitempty"`
}

func (a AuthError) Error() string {
	return fmt.Sprintf("%v (%v) %v", a.Type, a.Message, a.Cause)
}

func Authenticate(acct *Account, pwd string) error {
	l := login{
		Agent: struct {
			Name    string `json:"name"`
			Version int    `json:"version"`
		}{
			Name:    "Minecraft",
			Version: 1,
		},
		Username: acct.Login,
		Secret:   pwd,
		ClientID: acct.ClientToken,
	}
	//l := newMinecraftLogin(acct.Login, pwd, acct.ClientToken)
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

	if resp.StatusCode == http.StatusOK {
		if json.Unmarshal(body, acct) != nil {
			return err
		}

		return nil
	} else {
		var authErr AuthError
		if json.Unmarshal(body, &authErr) != nil {
			return err
		}

		return authErr
	}

}

func Validate(acct Account) (bool, error) {
	t := struct {
		AccessToken string `json:"accessToken"`
		ClientToken string `json:"clientToken"`
	}{
		AccessToken: acct.AccessToken,
		ClientToken: acct.ClientToken,
	}

	body, err := json.MarshalIndent(t, "", "    ")
	if err != nil {
		return false, err
	}
	req, err := http.NewRequest("POST", "https://authserver.mojang.com/validate", bytes.NewBuffer(body))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	if resp.StatusCode == http.StatusNoContent {
		return true, nil
	} else if resp.StatusCode == http.StatusForbidden {
		return false, nil
	} else {
		return false, errors.New(resp.Status)
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
