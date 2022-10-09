package services

import (
	"github.com/golang-jwt/jwt"
	"github.com/quanndh/go-app/public/resources"
	"time"
)

type UserClaims struct {
	*jwt.StandardClaims
	UserId uint `json:"userId"`
}

type JwtService struct {
}

func NewJwtService() IJwtService {
	return &JwtService{}
}

func (j JwtService) Generate(payload *resources.UserResource) (string, error) {
	privateKey := []byte("key")

	token := jwt.New(jwt.SigningMethodHS256)

	claims := token.Claims.(jwt.MapClaims)
	claims["userId"] = payload.ID
	claims["exp"] = time.Now().Add(720 * time.Hour).Unix()
	claims["iat"] = time.Now().Unix()

	tokenString, err := token.SignedString(privateKey)

	return tokenString, err
}

func (j JwtService) Verify(tokenString string) (*UserClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &UserClaims{}, func(token *jwt.Token) (interface{}, error) {

		return []byte("key"), nil
	})

	if err != nil {
		return nil, err
	}

	claims := token.Claims.(*UserClaims)
	
	return claims, nil
}
