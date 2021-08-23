package main

type SoundcloudUrl struct {
	Url string `json:"url"`
}

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
	BearerToken string `json:"bearerToken"`
}

func (u User) Read(p []byte) (n int, err error) {
	return 
	panic("implement me")
}

type BearerToken struct {
	BearerToken string `json:"bearerToken"`
}

type AuthenticatedSoundcloudUrl struct {
	Url string `json:"url"`
	BearerToken string `json:"bearerToken"`
}