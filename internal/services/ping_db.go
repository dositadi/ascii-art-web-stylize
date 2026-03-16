package services

import m "acad.learn2earn.ng/git/dositadi/ascii-art-web-stylize/pkg/models"

func (s *Service) CheckDBHealth() *m.Error {
	if err := s.Repository.PingDB(); err != nil {
		return err
	}
	return nil
}
