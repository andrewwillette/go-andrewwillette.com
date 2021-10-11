package main

type SoundcloudUrl struct {
	Url string `json:"url"`
}

type User struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	BearerToken string `json:"bearerToken"`
}

type BearerToken struct {
	BearerToken string `json:"bearerToken"`
}

type AuthenticatedSoundcloudUrl struct {
	Url         string `json:"url"`
	BearerToken string `json:"bearerToken"`
}
