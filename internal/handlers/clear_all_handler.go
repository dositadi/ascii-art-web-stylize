package handlers

import "net/http"

func (h *Handler) ClearAllHandler(w http.ResponseWriter, r *http.Request) {
	user_id := h.GetUserID(r)

	err := h.Service.ClearAllSavedAscii(r.Context(), user_id)
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
		return
	}

	err2 := h.Service.RenderHistoryPage(w, r, "")
	if err2 != nil {
		http.Error(w, err2.Details, http.StatusInternalServerError)
		return
	}
}
