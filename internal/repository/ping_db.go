package repository

import m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"

func (r *ServiceRepo) PingDB() *m.Error {
	err := r.DB.Ping()
	if err != nil {
		return &m.Error{
			Error:   "Ping error.",
			Details: err.Error(),
			Code:    "500",
		}
	}
	return nil
}
