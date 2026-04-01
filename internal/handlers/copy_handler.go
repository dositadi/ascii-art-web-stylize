package handlers

import "net/http"

func (h *Handler) CopyHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Service.CopyAscii(w, r)
	if err != nil {
		http.Error(w, "Could not copy", http.StatusInternalServerError)
	}
}
