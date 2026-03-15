package utils

import (
	"time"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	"github.com/golang-jwt/jwt/v5"
)

func GenerateAccessJWT(user m.ActiveUser) (string, *m.Error) {
	claims := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.RegisteredClaims{
		ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(5 * time.Hour))),
	})

	//claims.Method.Sign("", user)

	jwt, err := claims.SignedString("")
	if err != nil {
		return "", nil
	}

	return jwt, nil
}
