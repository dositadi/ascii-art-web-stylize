package utils

import (
	"golang.org/x/crypto/bcrypt"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

func HashPassword(password string) (string, *m.Error) {
	hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", &m.Error{
			Error:   SERVER_ERR,
			Details: SERVER_ERR_DETAIL,
			Code:    SERVER_ERR_CODE,
		}
	}

	return string(hashedPass), nil
}

func ComparePasswordAndHash(password string) *m.Error {
	hashedPass, err := HashPassword(password)
	if err != nil {
		return &m.Error{
			Error:   NOT_FOUND_ERR,
			Details: NOT_FOUND_DETAIL,
			Code:    NOT_FOUND_CODE,
		}
	}

	if err2 := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(password)); err2 != nil {
		return &m.Error{
			Error:   NOT_FOUND_ERR,
			Details: NOT_FOUND_DETAIL,
			Code:    NOT_FOUND_CODE,
		}
	}
	return nil
}
