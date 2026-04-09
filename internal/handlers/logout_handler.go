package handlers

import (
	"net/http"
	"time"

	u "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (h *Handler) LogoutHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Authorization", "")
	http.SetCookie(w, &http.Cookie{Name: "access_token", Value: "", HttpOnly: true, Secure: false, SameSite: http.SameSiteLaxMode, Path: "/", Expires: time.Now().Add(time.Duration(5 * time.Microsecond))})
	http.Redirect(w, r, u.WELCOME_ROUTE, http.StatusSeeOther)
}
