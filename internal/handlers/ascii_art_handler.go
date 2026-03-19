package handlers

import (
	"net/http"

	it "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/web/templates/inner_templates"
)

func (s *Handler) AsciiArtPageHandler(w http.ResponseWriter, r *http.Request) {
	err := it.ASCIITemplate(w, "")
	if err != nil {
		http.Error(w, err.Details, http.StatusInternalServerError)
	}
}
