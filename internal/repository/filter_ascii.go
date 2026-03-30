package repository

import (
	"context"
	"database/sql"
	"errors"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (r *ServiceRepo) Filter(ctx context.Context, font, user_id string) ([]m.Ascii, *m.Error) {
	rows, err := r.DB.QueryContext(ctx, h.FILTER_ASCII, user_id, font)
	if err != nil {
		return nil, &m.Error{
			Error:   h.SERVER_ERR,
			Details: err.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	var asciiArts []m.Ascii

	for rows.Next() {
		var asciiArt m.Ascii

		if err2 := rows.Scan(&asciiArt.Id, &asciiArt.InputText, &asciiArt.Font, &asciiArt.AsciiText); err2 != nil {
			if errors.Is(err2, sql.ErrConnDone) {
				return nil, &m.Error{
					Error:   h.CONN_LOST_ERR,
					Details: h.CONN_LOST_ERR_DETAIL,
					Code:    h.CONN_LOST_ERR_CODE,
				}
			}
			return nil, &m.Error{
				Error:   h.SERVER_ERR,
				Details: err2.Error(),
				Code:    h.SERVER_ERR,
			}
		}
		asciiArts = append(asciiArts, asciiArt)
	}

	if err3 := rows.Err(); err != nil {
		return nil, &m.Error{
			Error:   h.SERVER_ERR,
			Details: err3.Error(),
			Code:    h.SERVER_ERR,
		}
	}
	return asciiArts, nil
}
