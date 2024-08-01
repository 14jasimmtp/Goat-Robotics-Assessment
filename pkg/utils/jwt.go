package utils

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/14jasimmtp/Goat-Robotics-Assessment/pkg/db"
	"github.com/golang-jwt/jwt/v4"
	"github.com/spf13/viper"
)

type Utils struct {
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
	TokenString, err := token.SignedString([]byte(viper.GetString("ATokenSecret")))
	if err != nil {
		return "", err
	}
	return TokenString, nil
}

func IsValidAccessToken(secretKey, tokenString string) (*ClientClaims, error) {

	token, err := jwt.ParseWithClaims(tokenString, &ClientClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	})

	if err != nil || !token.Valid {
		log.Println("Error occurred while parsing token:", err)
		return nil, errors.New(`error parsing token. Token not valid`)
	}

	if claims, ok := token.Claims.(*ClientClaims); ok && token.Valid {

		if claims.ExpiresAt.Before(time.Now()) {
			fmt.Println("token expired")
			return nil, errors.New(`token expired`)
		}

		return claims, nil

	} else {
		fmt.Println("Error occurred while decoding token:", err)
		return nil, errors.New(`error in decoding token`)
	}
}
