package service

import (
	"fmt"
	"time"
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

func (s *Service) RedirectUrl(url string) (string, error) {

	urlData, err := s.st.Cache.Get(url)
	if err == nil {
		fmt.Println("From Cached Data")
	} else {
		fmt.Println("Getting data fron original database.....")
		urlData, err = s.st.RedirectUrl(url)
		if err != nil {
			return "", err
		}
		var ttl time.Duration
		if !urlData.ExpiryAt.IsZero() {
			ttl = time.Until(urlData.ExpiryAt)
		} else {
			ttl = time.Hour
		}
		_ = s.st.Cache.Set(url, *urlData, ttl)
		fmt.Println("Inserted into Cache from DB")
	}

	if !urlData.ExpiryAt.IsZero() && time.Now().After(urlData.ExpiryAt) {
		return "", fmt.Errorf("URL Expired")
	}
	if urlData.MaxClickLimit > 0 {
		err = s.st.Cache.CountClick(url, urlData.MaxClickLimit, 2*time.Minute)
		if err != nil {
			return "", err
		}
	}

	err = s.st.IncrementClickCount(urlData.Alias)
	if err != nil {
		return "", fmt.Errorf("failed to Update Click Count %v", err)
	}
	return urlData.Original, nil
}
