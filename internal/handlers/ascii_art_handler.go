package handlers

import (
	"fmt"
	"net/http"
)

func (s *Handler) AsciiArtPageHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Println("ascii: ", r.Context().Value("user_id"))
	err := s.Service.RenderAsciiArtPage(w, r)
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
	}
}
