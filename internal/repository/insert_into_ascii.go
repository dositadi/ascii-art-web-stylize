package repository

import (
	"context"
	"database/sql"
	"errors"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (r *ServiceRepo) InsertAscii(ctx context.Context, ascii m.Ascii, user_id string) *m.Error {
	tx, err := r.DB.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		if errors.Is(err, sql.ErrTxDone) {
			return &m.Error{
				Error:   h.SERVER_ERR,
				Details: err.Error(),
				Code:    h.SERVER_ERR_CODE,
			}
		}
	}

	exists, err2 := r.CheckIfAsciiExists(ctx, ascii.Id, user_id)
	if err2 != nil {
		return err2
	}

	if exists {
		return &m.Error{
			Error:   h.CONFLICT_ERR,
			Details: "Ascii has been saved already.",
			Code:    h.CONFLICT_ERR_CODE,
		}
	}

	_, err3 := tx.ExecContext(ctx, h.INSERT_INTO_ASCII_TEXTS, ascii.Id, ascii.UserId, ascii.InputText, ascii.Font, ascii.AsciiText)
	if err3 != nil {
		if err4 := tx.Rollback(); err4 != nil {
			return &m.Error{
				Error:   h.SERVER_ERR,
				Details: err4.Error(),
				Code:    h.SERVER_ERR_CODE,
			}
		}
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err3.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}

	if err5 := tx.Commit(); err5 != nil {
		return &m.Error{
			Error:   h.SERVER_ERR,
			Details: err5.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}
	return nil
}
