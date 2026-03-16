package models

import "github.com/golang-jwt/jwt/v5"

type User struct {
	Id             string `json:"id"`
	Name           string `json:"name"`
	Email          string `json:"email"`
	HashedPassword string `json:"hashed_password"`
	CreatedAt      string `json:"created_at"`
}

type ActiveUser struct {
	Id    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	jwt.RegisteredClaims
}
