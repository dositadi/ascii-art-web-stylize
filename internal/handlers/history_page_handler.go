package handlers

import "net/http"

func (s *Handler) HistoryPageHandler(w http.ResponseWriter, r *http.Request) {
	err := s.Service.RenderHistoryPage(w, r, "")
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
		return
	}
}
