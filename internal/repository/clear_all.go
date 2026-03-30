package repository

import (
	"context"
	"database/sql"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (r *ServiceRepo) ClearAll(ctx context.Context, user_id string) *m.Error {
	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	_, err2 := tx.ExecContext(ctx, h.CLEAR_ALL_USER_DATA, user_id)
	if err2 != nil {
		if err3 := tx.Rollback(); err3 != nil {
			return &m.Error{
				Error:   h.SERVER_ERR,
				Details: err3.Error(),
				Code:    h.SERVER_ERR_CODE,
			}
		}
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err2.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	if err4 := tx.Commit(); err4 != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err4.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}
	return nil
}
