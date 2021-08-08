package models

type SoundcloudUrl struct {
	Url string `json:"url"`
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type BearerToken struct {
	BearerToken string `json:"bearerToken"`
}