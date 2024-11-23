package utils

import (
	"errors"
	"time"
	"github.com/golang-jwt/jwt/v5"
)

var secretKey string = "MySecretKey"

func GenerateToken(id int64, email string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id": id,
		"email": email,
		"exp": time.Now().Add(time.Hour * 24).Unix(),
	})
	return token.SignedString([]byte(secretKey))
}

func VerifyToken(token string) (int64, error) {
	parsedToken, err := jwt.Parse(token , func(token *jwt.Token)(interface{}, error){
		_, ok := token.Method.(*jwt.SigningMethodHMAC)
		if !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(secretKey), nil
	})

	if err != nil {
        return 0, errors.New("could not parse token")
    }

	tokenIsValid := parsedToken.Valid

	if !tokenIsValid{
		return 0, errors.New("token is invalid")
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)

	if !ok {
		return 0, errors.New("Unexpected claims type")
	}

	userId := int64(claims["id"].(float64))
	return userId, nil
}