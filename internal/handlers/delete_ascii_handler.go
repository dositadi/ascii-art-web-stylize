package handlers

import (
	"net/http"

	u "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (h *Handler) DeleteAsciiHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	err := h.Service.DeleteAscii(r.Context(), id)
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
		return
	}

	w.Header().Set("HX-Redirect", u.HISTORY_ROUTE+u.ALL_HISTORY_QUERY)
}
