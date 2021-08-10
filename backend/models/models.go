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

type AuthenticatedSoundcloudUrl struct {
	Url string `json:"url"`
	BearerToken string `json:"bearerToken"`
}