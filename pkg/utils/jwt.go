package utils

import (
	"time"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/db"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type Utils struct{

}

type ClientClaims struct {
	ID    uint   `jsom:"id"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}

func GenerateAccessToken(user *db.Users) (string, error) {
	claims := ClientClaims{
		ID:    user.ID,
		Email: user.Email,
		RegisteredClaims: jwt.RegisteredClaims{
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			Issuer:    "Clockify",
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	TokenString, err := token.SignedString([]byte(viper.GetString("SecretKey")))
	if err != nil {
		return "", err
	}
	return TokenString, nil
}
