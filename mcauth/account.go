package mcauth

type Profile struct {
	Id       string     `json:"id"`
	PlayerName string     `json:"name"`
	Legacy     bool       `json:"legacy,omitempty"`
}

type Account struct {

	Login string

	AccessToken   string `json:"accessToken"`
	ClientToken   string `json:"clientToken"`
	Authenticated bool   `json:"-" sql:"-"`

	AvailableProfiles []Profile `json:"availableProfiles" sql:"-"`
	Profile   `json:"selectedProfile" sql:"profile"`
}
