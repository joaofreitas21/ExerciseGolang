package models

//import //"github.com/golang-jwt/jwt"
import "github.com/dgrijalva/jwt-go"

type Claim struct {
	ID uint32 `json:"id"`
	jwt.StandardClaims
}
