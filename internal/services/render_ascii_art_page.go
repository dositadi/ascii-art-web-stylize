package services

import (
	"html/template"
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (s *Service) RenderSessionExpiredPage(w http.ResponseWriter, r *http.Request) *m.Error {
	temp, err := template.New("session_expired.html").ParseFiles("web/static/auth_pages/session_expired.html", "web/templates/session_expired_partial.html")
	if err != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	sessionExpiredPageData := struct {
		LoginPageRoute string
	}{
		LoginPageRoute: h.LOGIN_ROUTE,
	}

	if s.GetHxRequestStatus(r) {
		if err2 := temp.ExecuteTemplate(w, "session-expired", sessionExpiredPageData); err2 != nil {
			return &m.Error{
				Error:   h.SERVER_ERR,
				Details: err2.Error(),
				Code:    h.SERVER_ERR_CODE,
			}
		}
	} else {
		if err3 := temp.Execute(w, sessionExpiredPageData); err3 != nil {
			return &m.Error{
				Error:   h.SERVER_ERR,
				Details: err3.Error(),
				Code:    h.SERVER_ERR_CODE,
			}
		}
	}
	return nil
}
