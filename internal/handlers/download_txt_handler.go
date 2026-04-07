package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) DownloadAsciiTxtHandler(w http.ResponseWriter, r *http.Request) {
	text := r.URL.Query().Get("text")
	font := r.URL.Query().Get("font")

	fmt.Println(text, font)
	err := h.Service.DownloadAsTxt(w, text, font)
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
		return
	}
}
