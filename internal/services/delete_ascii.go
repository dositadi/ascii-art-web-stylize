package services

import (
	"context"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

func (s *Service) DeleteAscii(ctx context.Context, id string) *m.Error {
	err := s.Repository.DeleteFromAscii(ctx, id)
	if err != nil {
		return &m.Error{
			Error:   err.Error,
			Details: err.Details,
			Code:    err.Code,
		}
	}
	return nil
}
