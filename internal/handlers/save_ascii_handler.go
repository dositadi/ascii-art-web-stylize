package handlers

import (
	"fmt"
	"net/http"

	u "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (h *Handler) SaveAsciiHandler(w http.ResponseWriter, r *http.Request) {
	val := r.Context().Value("user_id")
	var user_id string

	if id, ok := val.(string); ok {
		user_id = id
	}

	text := r.FormValue(u.TEXT_KEY)
	banner := r.FormValue(u.BANNER_KEY)

	if text == "" {
		fmt.Fprintf(w, `<div>`)
		return
	}

	err := h.Service.SaveAscii(r.Context(), text, banner, user_id)
	if err != nil {
		http.Error(w, "Save failed", http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", u.HISTORY_ROUTE+u.ALL_HISTORY_QUERY)
}
