package service

import (
	"url-shortener/models"
	"url-shortener/store"
)

type Service struct {
	st *store.Store
}

func NewService(s *store.Store) *Service {
	return &Service{
		st: s,
	}
}

func (s *Service) ShortenUrl(urlReq *models.URLData) error {

	return s.st.ShortenUrl(urlReq)
}

func (s *Service) RedirectUrl(url string) string {
	return s.st.RedirectUrl(url)
}
