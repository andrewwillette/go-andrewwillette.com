package server

type SoundcloudUrlJson struct {
	Url     string `json:"url"`
	UiOrder int    `json:"uiOrder"`
	Id      int    `json:"id"`
}

type UserJson struct {
	Username    string `json:"username"`
	Password    string `json:"password"`
	BearerToken string `json:"bearerToken"`
}

type BearerTokenJson struct {
	BearerToken string `json:"bearerToken"`
}

type AuthenticatedSoundcloudUrlJson struct {
	Url         string `json:"url"`
	BearerToken string `json:"bearerToken"`
}
