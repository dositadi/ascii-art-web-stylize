package services

import (
	"fmt"
	"html/template"
	"net/http"
	"strconv"
	"time"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

var (
	count  int
	reduce int
	items  int
)

func (s *Service) RenderHistoryPage(w http.ResponseWriter, r *http.Request, font string) *m.Error {
	funcMap := template.FuncMap{
		"Format": func(t time.Time) string {
			return t.Format("02 Jan 2006 15:04")
		},
	}
	temp, err := template.New("history.html").Funcs(funcMap).ParseFiles("web/static/internal_pages/history.html", "web/templates/history_partial.html", "web/templates/history_tabs_partial.html")
	if err != nil {
		fmt.Println(err.Error())
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

	_, userName, _, err2 := s.Repository.GetHashedPasswordIDAndName(r.Context(), &user_id, nil)
	if err2 != nil {
		fmt.Println(err2.Details)
		return err2
	}

	namePrefix := s.GetNamePrefix(userName)

	var length int
	var err6 *m.Error

	if font == "" || font == "all" {
		length, err6 = s.Repository.GetTableLenght(r.Context(), user_id, "")
	} else {
		length, err6 = s.Repository.GetTableLenght(r.Context(), user_id, font)
	}

	if err6 != nil {
		fmt.Println(err6.Details)
		return err6
	}

	var next, prev int

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	limit := 2

	count = page
	count++
	next = count

	reduce = page
	reduce--
	prev = reduce

	offset := (page - 1) * limit

	var asciiArts []m.Ascii
	var err5 *m.Error

	if font == "" || font == "all" {
		asciiArts, err5 = s.Repository.GetAllUsersSavedAscii(r.Context(), user_id, limit, offset, "")
	} else {
		asciiArts, err5 = s.Repository.GetAllUsersSavedAscii(r.Context(), user_id, limit, offset, font)
	}

	if err5 != nil {
		fmt.Println("Entered here: ", err5.Details)
		return &m.Error{
			Error:   h.PAGE_PARSING_CODE,
			Details: err.Error(),
			Code:    h.SERVER_ERR,
		}
	}

	items = page * len(asciiArts)

	var disableNext, disablePrev bool

	if len(asciiArts) < limit || items == length {
		disableNext = true
	}

	if page == 1 {
		disablePrev = true
	}

	historyPageDetail := struct {
		DisableNext          bool
		DisablePrev          bool
		PageRoute            string
		PrevPageRoute        string
		History              string
		AsciiArts            []m.Ascii
		UserName             string
		NamePrefix           string
		AsciiRoute           string
		AboutRoute           string
		HelpRoute            string
		ContributorsRoute    string
		DeleteRoute          string
		StandardFilterRoute  string
		ShadowFilterRoute    string
		TinkertoyFilterRoute string
		AllFilterRoute       string
		ClearAllRoute        string
	}{
		DisableNext:          disableNext,
		DisablePrev:          disablePrev,
		PageRoute:            h.HISTORY_ROUTE + "?font=all&page=" + strconv.Itoa(next),
		PrevPageRoute:        h.HISTORY_ROUTE + "?font=all&page=" + strconv.Itoa(prev),
		AsciiArts:            asciiArts,
		UserName:             userName,
		NamePrefix:           namePrefix,
		AsciiRoute:           h.ASCII_ROUTE,
		AboutRoute:           h.ABOUT_US_ROUTE,
		HelpRoute:            h.HELP_ROUTE,
		ContributorsRoute:    h.CONTRIBUTORS_ROUTE,
		DeleteRoute:          h.DELETE_ROUTE,
		StandardFilterRoute:  h.HISTORY_ROUTE + h.STANDARD_HISTORY_QUERY,
		ShadowFilterRoute:    h.HISTORY_ROUTE + h.SHADOW_HISTORY_QUERY,
		TinkertoyFilterRoute: h.HISTORY_ROUTE + h.TINKERTOY_HISTORY_QUERY,
		AllFilterRoute:       h.HISTORY_ROUTE + h.ALL_HISTORY_QUERY,
		ClearAllRoute:        h.CLEAR_ALL_ROUTE,
	}

	if s.GetHxRequestStatus(r) {
		if err3 := temp.ExecuteTemplate(w, "history", historyPageDetail); err3 != nil {
			fmt.Println(err3.Error())
			return &m.Error{
				Error:   h.PAGE_PARSING_CODE,
				Details: err3.Error(),
				Code:    h.SERVER_ERR,
			}
		}
	} else {
		if err4 := temp.Execute(w, historyPageDetail); err4 != nil {
			fmt.Println(err4.Error())
			return &m.Error{
				Error:   h.PAGE_PARSING_CODE,
				Details: err4.Error(),
				Code:    h.SERVER_ERR,
			}
		}
	}
	return nil
}
