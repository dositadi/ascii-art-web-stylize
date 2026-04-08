package services

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	"github.com/google/uuid"
)

func (s *Service) DownloadAsTxt(w http.ResponseWriter, text, font string) *m.Error {
	formattedAscii, err := s.FormatAscii(text, font)
	fmt.Println(font, text)
	if err != nil {
		return err
	}

	downloadId := uuid.NewString()

	attachment := fmt.Sprintf("attachment; filename=ascii_forge_%s_%s.txt", text, downloadId)

	w.Header().Set("Content-Disposition", attachment)
	w.Header().Set("Content-Type", "text/plain")
	w.Header().Set("Content-Length", strconv.Itoa(len(formattedAscii)))
	//w.Header().Set("HX-Reswap", "none")

	_, err2 := io.Copy(w, strings.NewReader(formattedAscii))
	if err2 != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err2.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}
	return nil
}
