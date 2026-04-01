package services

import (
	"context"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	"github.com/google/uuid"
)

func (s *Service) SaveAscii(ctx context.Context, text, banner, user_id string) *m.Error {
	formattedAsciiWords, err := s.FormatAscii(text, banner)
	if err != nil {
		return err
	}

	ascii := m.Ascii{
		Id:        uuid.NewString(),
		UserId:    user_id,
		InputText: text,
		Font:      banner,
		AsciiText: formattedAsciiWords,
	}

	err3 := s.Repository.InsertAscii(ctx, ascii)
	if err3 != nil {
		return err3
	}
	return nil
}
