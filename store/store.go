package store

import (
	"database/sql"
	"fmt"
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
	query := `INSERT into urldata(original,shortend) VALUES($1,$2)`
	_, err := s.db.Exec(query, urlData.Original, urlData.Shortend)
	if err != nil {
		fmt.Printf("Error inserting into DB: %v\n", err)
		return "", err
	}
	return DefaultDomain + urlData.Shortend, nil
}

func (s *Store) RedirectUrl(url string) string {
	var urlData models.URLData
	query := `SELECT original from urlData where shortend = $1`
	err := s.db.QueryRow(query, url).Scan(&urlData.Original)
	if err != nil {
		fmt.Printf("Error Fetching the original url: %v\n", err)
		return ""
	}
	return urlData.Original
}

func (s *Store) SearchAliasExsists(urlData *models.URLData) (bool, error) {
	var existingalias string
	query := `SELECT shortend from urlData where shortend = $1`
	err := s.db.QueryRow(query, urlData.Shortend).Scan(&existingalias)
	if err == sql.ErrNoRows {
		return false, nil
	}
	return true, nil
}
