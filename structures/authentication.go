package structures

import "github.com/dgrijalva/jwt-go"

type Credentials struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type Claims struct {
	Id_user int `json:"id_user"`
	jwt.StandardClaims
}

type Token struct {
	Auth string `json:"auth"`
}
