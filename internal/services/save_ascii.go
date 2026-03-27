package services

import (
	"context"
	"fmt"

	m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"
	"github.com/google/uuid"
)

func (s *Service) SaveAscii(ctx context.Context, text, banner, user_id string) *m.Error {
	latinWords, err := s.AsciiTransformer.SplitInputByNewline(text)
	if err != nil {
		fmt.Println("Entered 8")
		return err
	}

	asciiWords, _, err2 := s.AsciiTransformer.ReadWords(latinWords, banner)
	if err2 != nil {
		fmt.Println("Entered 7")
		return err2
	}

	formattedAsciiWords := s.AsciiTransformer.FormatAsciiWords(asciiWords)

	ascii := m.Ascii{
		Id:        uuid.NewString(),
		UserId:    user_id,
		InputText: text,
		Font:      banner,
		AsciiText: formattedAsciiWords,
	}
	fmt.Println(user_id)

	err3 := s.Repository.InsertAscii(ctx, ascii, user_id)
	if err3 != nil {
		return err3
	}
	return nil
}
