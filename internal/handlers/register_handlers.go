package handlers

import (
	"fmt"
	"net/http"
	"strings"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	at "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/web/templates/auth_templates"
)

func (s *Handler) RegisterHandler(w http.ResponseWriter, r *http.Request) {
	first_name := strings.TrimSpace(r.FormValue("name"))
	last_name := r.FormValue("last_name")
	user_email := strings.TrimSpace(r.FormValue("email"))
	user_password := strings.TrimSpace(r.FormValue("password"))

	user_name := first_name + " " + last_name

	fmt.Println(user_email)
	fmt.Printf("Value: %q | Length: %d\n", user_name, len(user_name))

	if user_email == "" {
		err := m.Error{Error: h.EMPTY_EMAIL_FIELD, Details: h.EMPTY_EMAIL_FIELD_DETAIL, Code: h.EMPTY_EMAIL_FIELD_CODE}
		h.ErrorResponse(w, h.ErrorToJson(err), http.StatusBadRequest)
		return
	}
	if user_name == "" {
		err := m.Error{Error: h.EMPTY_NAME_FIELD, Details: h.EMPTY_NAME_FIELD_DETAIL, Code: h.EMPTY_NAME_FIELD_CODE}
		h.ErrorResponse(w, h.ErrorToJson(err), http.StatusBadRequest)
		return
	}
	if len(user_password) < 8 || user_password == "" {
		err := m.Error{Error: h.EMPTY_PASSWORD_FIELD, Details: h.EMPTY_PASSWORD_FIELD_DETAIL, Code: h.EMPTY_PASSWORD_FIELD_CODE}
		h.ErrorResponse(w, h.ErrorToJson(err), http.StatusBadRequest)
		return
	}
	if !h.IsEmail(user_email) {
		err := m.Error{Error: h.BAD_EMAIL_FORMAT, Details: h.BAD_EMAIL_FORMAT_DETAIL, Code: h.BAD_EMAIL_FORMAT_CODE}
		h.ErrorResponse(w, h.ErrorToJson(err), http.StatusBadRequest)
		return
	}

	fmt.Println(user_name)

	user := m.User{
		Name:  user_name,
		Email: user_email,
	}

	ctx := r.Context()

	err2 := s.Service.RegisterUser(ctx, &user, user_password)
	if err2 != nil {
		if err2.Error == h.CONFLICT_ERR {
			err := h.ErrorToJson(*err2)
			h.ErrorResponse(w, err, http.StatusConflict)
		} else if err2.Error == h.SERVER_ERR {
			err := h.ErrorToJson(*err2)
			h.ErrorResponse(w, err, http.StatusInternalServerError)
		} else {
			err := h.ErrorToJson(*err2)
			h.ErrorResponse(w, err, http.StatusBadRequest)
		}
	}

	http.Redirect(w, r, "/auth/login", http.StatusSeeOther)
}

func (s *Handler) RegisterPageHandler(w http.ResponseWriter, r *http.Request) {
	err := at.SignUpPageTemplate(w)
	if err != nil {
		err := h.ErrorToJson(m.Error{Error: err.Error, Details: err.Details, Code: err.Code})
		h.ErrorResponse(w, err, http.StatusBadRequest)
		return
	}
}
