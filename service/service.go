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
	if urlReq.Alias == "" {
		//return "", fmt.Errorf("alias can't be empty")
		urlReq.Alias = utils.GenerateRandomAlias()

	}
	exsists, err := s.st.SearchAliasExsists(urlReq)
	if exsists {
		fmt.Println(" exsis its true")
		return "", fmt.Errorf("alias already in use")

	}
	if err != nil {
		fmt.Println("exsis its error")
		return "", err
	}
	return s.st.ShortenUrl(urlReq)

}

func (s *Service) RedirectUrl(url string) (string,error) {
	return s.st.RedirectUrl(url)
}
