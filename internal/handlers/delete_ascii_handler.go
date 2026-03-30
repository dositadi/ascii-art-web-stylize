package handlers

import (
	"fmt"
	"net/http"
)

func (h *Handler) DeleteAsciiHandler(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")

	err := h.Service.DeleteAscii(r.Context(), id)
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	err2 := h.Service.RenderHistoryPage(w, r)
	if err != nil {
		http.Error(w, err2.Details, http.StatusInternalServerError)
		return
	}
}
