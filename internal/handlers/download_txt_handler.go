package handlers

import (
	"net/http"

	u "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (h *Handler) DownloadAsciiTxtHandler(w http.ResponseWriter, r *http.Request) {
	text := r.FormValue(u.TEXT_KEY)
	font := r.FormValue(u.BANNER_KEY)

	err := h.Service.DownloadAsTxt(w, text, font)
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
		return
	}
}
