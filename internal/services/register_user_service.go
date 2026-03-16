package services

import (
	"context"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	h "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/utils"
	"github.com/google/uuid"
)

func (s *Service) RegisterUser(ctx context.Context, user *m.User, password string) *m.Error {
	hashedPassword, err := h.HashPassword(password)
	if err != nil {
		return err
	}

	user.HashedPassword = hashedPassword
	user.Id = uuid.NewString()

	err1 := s.Repository.InsertUser(ctx, *user)
	if err1 != nil {
		return err
	}
	return nil
}
