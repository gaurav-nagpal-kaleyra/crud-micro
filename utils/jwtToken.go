package utils

import (
	"crud-micro/constant"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
	"go.uber.org/zap"
)

func CreateJwtToken(userId string) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"userId": userId,
		"exp":    time.Now().Add(time.Hour * 24).Unix(),
	})

	signedToken, err := token.SignedString([]byte(os.Getenv(constant.SECRETKEY)))
	if err != nil {
		zap.L().Error("error creating token", zap.Error(err))
		return "", err
	}

	return signedToken, nil
}
