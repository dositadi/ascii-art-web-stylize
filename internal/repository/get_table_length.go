package repository

import (
	"context"
	"fmt"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
)

func (r *ServiceRepo) GetTableLenght(ctx context.Context) (int, *m.Error) {
	row := r.DB.QueryRowContext(ctx, h.GET_TABLE_LENGHT)

	var length int

	if err := row.Scan(&length); err != nil {
		fmt.Println(err.Error())
		return 0, &m.Error{
			Error:   h.SERVER_ERR,
			Details: err.Error(),
			Code:    h.SERVER_ERR_CODE,
		}
	}
	return length, nil
}
