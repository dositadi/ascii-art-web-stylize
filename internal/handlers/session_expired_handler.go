package handlers

import "net/http"

func (h *Handler) SessionExpiredHandler(w http.ResponseWriter, r *http.Request) {
	err := h.Service.RenderSessionExpiredPage(w, r)
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
		return
	}
}
