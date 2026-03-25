package services

import (
	"html/template"
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *Service) RenderHistoryPage(w http.ResponseWriter, r *http.Request) *m.Error {
	temp, err := template.New("history.html").ParseFiles("web/static/internal_pages/history.html", "web/templates/history_partial.html")
	if err != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_CODE,
			Details: err.Error(),
			Code:    h.SERVER_ERR,
		}
	}

	historyPageDetail := struct {
		AsciiRoute string
		AboutRoute string
		HelpRoute  string
	}{
		AsciiRoute: h.ASCII_ROUTE,
	}

	if s.GetHxRequestStatus(r) {
		if err2 := temp.ExecuteTemplate(w, "history", historyPageDetail); err2 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_CODE,
				Details: err2.Error(),
				Code:    h.SERVER_ERR,
			}
		}
	} else {
		if err3 := temp.Execute(w, historyPageDetail); err3 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_CODE,
				Details: err3.Error(),
				Code:    h.SERVER_ERR,
			}
		}
	}
	return nil
}
