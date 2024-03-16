package login_user

type UserDataIn struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type UserDataOut struct {
	AccessToken string `json:"access_token"`
}
