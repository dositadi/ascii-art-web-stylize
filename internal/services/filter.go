package services

import (
	"net/http"
	"text/template"
	"time"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *Service) FilterAscii(w http.ResponseWriter, r *http.Request, font string) *m.Error {
	funcMap := template.FuncMap{
		"Format": func(t time.Time) string {
			return t.Format("02 Jan 2006 15:04")
		},
	}
	temp, err := template.New("history_tabs_partial.html").Funcs(funcMap).ParseFiles("web/templates/history_tabs_partial.html")
	if err != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_CODE,
			Details: err.Error(),
			Code:    h.SERVER_ERR,
		}
	}

	val := r.Context().Value("user_id")
	var user_id string

	if id, ok := val.(string); ok {
		user_id = id
	}

	var asciiArts []m.Ascii
	var err5 *m.Error

	if font != "all" {
		asciiArts, err5 = s.Repository.Filter(r.Context(), font, user_id)
		if err5 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_CODE,
				Details: err.Error(),
				Code:    h.SERVER_ERR,
			}
		}
	} else {
		asciiArts, err5 = s.Repository.GetAllUsersSavedAscii(r.Context(), user_id, 1, 1)
		if err5 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_CODE,
				Details: err.Error(),
				Code:    h.SERVER_ERR,
			}
		}
	}

	historyTabDetail := struct {
		AsciiArts   []m.Ascii
		AsciiRoute  string
		DeleteRoute string
	}{
		AsciiArts:   asciiArts,
		AsciiRoute:  h.ASCII_ROUTE,
		DeleteRoute: h.DELETE_ROUTE,
	}

	if s.GetHxRequestStatus(r) {
		if err3 := temp.ExecuteTemplate(w, "tabs", historyTabDetail); err3 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_CODE,
				Details: err3.Error(),
				Code:    h.SERVER_ERR,
			}
		}
	} else {
		if err4 := s.RenderHistoryPage(w, r); err4 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_CODE,
				Details: err4.Details,
				Code:    h.SERVER_ERR,
			}
		}
	}

	return nil
}
