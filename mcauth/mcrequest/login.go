package mcrequest

type Login struct{
	Agent `json:"agent"`
	Username string `json:"username"`
	Secret   string `json:"password"`
	ClientID string `json:"clientToken,omitempty"`
}

func NewMinecraftLogin(username, secret, clientID string) Login {

	return Login{
		Agent: Agent{Name:"Minecraft", Version: 1},
		Username: username,
		Secret: secret,
		ClientID: clientID,
	}
}