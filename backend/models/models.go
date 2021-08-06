package models

type SoundcloudUrl struct {
	Url string `json:"url"`
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type SessionKey struct {
	SessionKey string `json:"sessionKey"`
}