package service

import (
	"fmt"
	"url-shortener/models"
	"url-shortener/store"
	"url-shortener/utils"
)

type Service struct {
	st *store.Store
}

func NewService(s *store.Store) *Service {
	return &Service{
		st: s,
	}
}

func (s *Service) ShortenUrl(urlReq *models.URLData) (string, error) {
	if urlReq.Shortend == "" {
		//return "", fmt.Errorf("alias can't be empty")
		urlReq.Shortend = utils.GenerateRandomAlias()

	}
	exsists, err := s.st.SearchAliasExsists(urlReq)
	if exsists {
		fmt.Println(" exsisits true")
		return "", fmt.Errorf("alias already in use")

	}
	if err != nil {
		fmt.Println("exsisits error")
		return "", err
	}
	return s.st.ShortenUrl(urlReq)

}

func (s *Service) RedirectUrl(url string) string {
	return s.st.RedirectUrl(url)
}
