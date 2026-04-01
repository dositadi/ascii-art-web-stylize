package handlers

import (
	"net/http"
)

func (s *Handler) HistoryPageHandler(w http.ResponseWriter, r *http.Request) {
	font := r.URL.Query().Get("font")
	err := s.Service.RenderHistoryPage(w, r, font)
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
		return
	}
}
