package login_user

type UserDataIn struct {
	Username string `json:"username" example:"user"`
	Password string `json:"password" example:"password"`
}

type UserDataOut struct {
	AccessToken string `json:"access_token" example:"token"`
}
