package services

/* import (
	"net/http"
	"strconv"
	"text/template"
	"time"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

var (
	Count  int
	Reduce int
	Items  int
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

	_, userName, _, err2 := s.Repository.GetHashedPasswordIDAndName(r.Context(), &user_id, nil)
	if err2 != nil {
		return err2
	}

	namePrefix := s.GetNamePrefix(userName)

	page, _ := strconv.Atoi(r.URL.Query().Get("page"))

	limit := 2

	offset := (page - 1) * limit

	Count := page
	Count++
	next := Count

	Reduce := page
	Reduce--
	prev := Reduce

	var pageRoute string

	switch font {
	case "standard":
		pageRoute = h.STANDARD_FILTER_ROUTE
	case "shadow":
		pageRoute = h.SHADOW_FILTER_ROUTE
	case "tinkertoy":
		pageRoute = h.TINKERTOY_FILTER_ROUTE
	}

	var asciiArts []m.Ascii
	var err5 *m.Error

	totalAscii, err6 := s.Repository.GetTableLenght(r.Context())
	if err6 != nil {
		return err6
	}

	var disableNext, disablePrev bool

	if font != "all" {
		asciiArts, err5 = s.Repository.Filter(r.Context(), limit, offset, font, user_id)
		if err5 != nil {
			return &m.Error{
				Error:   h.PAGE_PARSING_CODE,
				Details: err.Error(),
				Code:    h.SERVER_ERR,
			}
		}

		items = len(asciiArts) * page

		if items == totalAscii || len(asciiArts) < limit {
			disableNext = true
		}

		if page == 1 {
			disablePrev = true
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
		DisableNext          bool
		DisablePrev          bool
		PageRoute            string
		PrevPageRoute        string
		Next                 int
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
		PageRoute:            pageRoute + "?page=" + strconv.Itoa(next),
		PrevPageRoute:        pageRoute + "?page=" + strconv.Itoa(prev),
		AsciiArts:            asciiArts,
		UserName:             userName,
		NamePrefix:           namePrefix,
		AsciiRoute:           h.ASCII_ROUTE,
		AboutRoute:           h.ABOUT_US_ROUTE,
		HelpRoute:            h.HELP_ROUTE,
		ContributorsRoute:    h.CONTRIBUTORS_ROUTE,
		DeleteRoute:          h.DELETE_ROUTE,
		StandardFilterRoute:  h.STANDARD_FILTER_ROUTE + h.HistoryQuery,
		ShadowFilterRoute:    h.SHADOW_FILTER_ROUTE + h.HistoryQuery,
		TinkertoyFilterRoute: h.TINKERTOY_FILTER_ROUTE + h.HistoryQuery,
		AllFilterRoute:       h.ALL_ASCII_FILTER_ROUTE + h.HistoryQuery,
		ClearAllRoute:        h.CLEAR_ALL_ROUTE,
	}

	if s.GetHxRequestStatus(r) {
		if err3 := temp.ExecuteTemplate(w, "history", historyTabDetail); err3 != nil {
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
*/
