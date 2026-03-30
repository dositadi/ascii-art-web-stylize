package services

import (
	"context"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
)

func (s *Service) ClearAllSavedAscii(ctx context.Context, user_id string) *m.Error {
	err := s.Repository.ClearAll(ctx, user_id)
	if err != nil {
		return err
	}
	return nil
}
