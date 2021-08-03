package main

type SoundcloudUrl struct {
	Url string `json:"url"`
}

type UserCredentials struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type LoginResponse struct {
	SessionKey string `json:"sessionKey"`
}