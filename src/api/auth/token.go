package auth

import (
	"errors"
	"net/http"
	"teste1/src/api/models"
	"teste1/src/api/responses"
	"teste1/src/config"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/dgrijalva/jwt-go/request"
	//"github.com/golang-jwt/jwt/request"
)
/*
	Func Gerar JWT Token
	Recebe o id do utilizador para guardar no token
*/

func GenerateToken(uid uint32) (string, error) {
	/*claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = uid
	claims["exp"] = time.Now().Add(time.Minute * 30).Unix()*/
	claims := models.Claim{
		ID: uid,
		StandardClaims: jwt.StandardClaims{
			Issuer:    "Joao Freitas",
			ExpiresAt: time.Now().Add(time.Minute * 30).Unix(),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claims)
	return token.SignedString(config.SECRETKEY)
}

/*
	Func: Decifrar o token
	Lidar com erros JWT
*/
func DecodeToken(w http.ResponseWriter, r *http.Request) *jwt.Token{
	token, err := request.ParseFromRequestWithClaims(
		r,
		request.OAuth2Extractor,
		&models.Claim{},
		func(t *jwt.Token) (interface{}, error) {
			return config.SECRETKEY, nil
		},
	)

	if err != nil {
		code := http.StatusUnauthorized
		switch err.(type) {
		case *jwt.ValidationError:
			vError := err.(*jwt.ValidationError)
			switch vError.Errors {
			case jwt.ValidationErrorExpired:
				err = errors.New("Your token has expired")
				responses.ERROR(w, code, err)
				return nil
			case jwt.ValidationErrorSignatureInvalid:
				err = errors.New("The signature is invalid")
				responses.ERROR(w, code, err)
				return nil
			default:
				responses.ERROR(w, code, err)
				return nil
			}
		}
	}

	return token

}
/*
func DecodeToken(r *http.Request) string{
	keys := r.URL.Query()
	token := keys.Get("token")
	if token != ""{
		return token
	}

	bearerToken := r.Header.Get("Authorization")
	if len(strings.Split(bearerToken," ")) == 2 {
		
		return strings.Split(bearerToken," ")[1]
	}
	return ""
}*/