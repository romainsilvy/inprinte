package structures

type GetAuthData struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Token struct {
	Auth string `json:"auth"`
}
