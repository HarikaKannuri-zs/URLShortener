package store

import (
	"database/sql"
	"fmt"
	"time"
	"url-shortener/models"

	_ "github.com/lib/pq"
)

type Store struct {
	db *sql.DB
}

const DefaultDomain = "https://mygourl/"

func NewStore(urlDb *sql.DB) *Store {
	return &Store{
		db: urlDb,
	}
}

func (s *Store) ShortenUrl(urlData *models.URLData) (string, error) {
	query := `INSERT into urldata(original,alias,expiry_at,max_click_limit) VALUES($1,$2,$3,$4)`
	_, err := s.db.Exec(query, urlData.Original, urlData.Alias, urlData.ExpiryAt, urlData.MaxClickLimit)
	if err != nil {
		fmt.Printf("Error inserting into DB: %v\n", err)
		return "", err
	}
	return DefaultDomain + urlData.Alias, nil
}

func (s *Store) RedirectUrl(aliasUurl string) (string, error) {
	var urlData models.URLData
	query := `SELECT original,expiry_at,click_cnt,max_click_limit from urlData where Alias = $1`
	err := s.db.QueryRow(query, aliasUurl).Scan(&urlData.Original, &urlData.ExpiryAt, &urlData.ClickCnt, &urlData.MaxClickLimit)
	if err != nil {
		return "", err
	}
	if !urlData.ExpiryAt.IsZero() && time.Now().After(urlData.ExpiryAt) {
		return "", fmt.Errorf("URL Expired")
	}
	if urlData.MaxClickLimit > 0 && urlData.ClickCnt >= urlData.MaxClickLimit {
		return "", fmt.Errorf("click Limit Reached. Wait for few minutes")
	}
	_, err = s.db.Exec(`UPDATE urlData SET click_cnt = click_cnt+1 where alias = $1`, aliasUurl)
	if err != nil {
		return "", fmt.Errorf("failed to Update Click Count %v", err)
	}

	return urlData.Original, nil
}

func (s *Store) SearchAliasExsists(urlData *models.URLData) (bool, error) {
	var existingalias string
	query := `SELECT alias from urlData where Alias = $1`
	err := s.db.QueryRow(query, urlData.Alias).Scan(&existingalias)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, nil
}
