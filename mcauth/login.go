package mcauth

type login struct {
	Agent    struct {
			 Name    string `json:"name"`
			 Version int    `json:"version"`
		 } `json:"agent"`
	Username string `json:"username"`
	Secret   string `json:"password"`
	ClientID string `json:"clientToken,omitempty"`
}

func newMinecraftLogin(username, secret, clientID string) login {

	return login{
		Agent: struct {
			Name    string `json:"name"`
			Version int `json:"version"`
		}{
			Name:    "Minecraft",
			Version: 1,
		},
		Username: username,
		Secret:   secret,
		ClientID: clientID,
	}
}
