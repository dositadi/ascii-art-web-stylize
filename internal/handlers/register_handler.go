package handlers

import (
	"net/http"

	"github.com/google/uuid"
)

/*
Id        string `json:"id"`
	Name      string `json:"name"`
	Email     string `json:"email"`
	Password  string `json:"password"`
	CreatedAt string `json:"created_at"`
*/

func (s *Handler) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	name := r.PostFormValue("name")
	email := r.PostFormValue("email")
	password := r.PostFormValue("password")
	id := uuid.NewString()

	if name == "" {
		
	}

}
