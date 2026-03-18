package innertemplates

import (
	"html/template"
	"net/http"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func ASCIITemplate(w http.ResponseWriter, output string) *m.Error {
	temp, err := template.New("ascii.html").ParseFiles("web/static/internal_pages/ascii.html")
	if err != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_ERROR,
			Details: err.Error(),
			Code:    h.PAGE_PARSING_CODE,
		}
	}

	asciiPageDetails := struct {
		LogOutRoute      string
		TextKey          string
		BannerKey        string
		Output           string
		DownloadImgRoute string
		DownloadTxtRoute string
		SaveOutputRoute  string
		ViewHistoryRoute string
	}{
		Output:           output,
		TextKey:          h.TEXT_KEY,
		BannerKey:        h.BANNER_KEY,
		LogOutRoute:      "",
		DownloadImgRoute: "",
		DownloadTxtRoute: "",
		SaveOutputRoute:  "",
		ViewHistoryRoute: "",
	}

	if err1 := temp.Execute(w, asciiPageDetails); err1 != nil {
		return &m.Error{
			Error:   h.PAGE_PARSING_ERROR,
			Details: err.Error(),
			Code:    h.PAGE_PARSING_CODE,
		}
	}
	return nil
}
