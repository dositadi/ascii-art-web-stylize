package repository

import (
	"context"
	"database/sql"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (r *ServiceRepo) DeleteFromAscii(ctx context.Context, id string) *m.Error {
	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err.Error(),
			Code:    h.SERVER_ERR_CODE,        
		}
	}

	exists, err1 := r.CheckIfAsciiExists(ctx, id)
	if err1 != nil {
		return err1
	}

	if !exists {
		return &m.Error{
			Error:   h.NOT_FOUND_ERR,
			Details: h.NOT_FOUND_DETAIL,
			Code:    h.NOT_FOUND_CODE,
		}
	}

	_, err2 := tx.ExecContext(ctx, h.DELETE_ASCII, id)
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
