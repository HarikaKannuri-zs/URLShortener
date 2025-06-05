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

func NewStore(urlDb *sql.DB) *Store {
	return &Store{
		db: urlDb,
	}
}

func (s *Store) ShortenUrl(urlData *models.URLData) error {

	query := `INSERT into urldata(original,shortend) VALUES($1,$2)`
	_, err := s.db.Exec(query, urlData.Original, urlData.Shortend)
	if err != nil {
		fmt.Printf("Error inserting into DB: %v\n", err)
		return err
	}
	return nil
}

func (s *Store) RedirectUrl(url string) string {
	var urlData models.URLData
	query := `SELECT original from urlData where shortend = $1`
	org := s.db.QueryRow(query, url).Scan(&urlData.Original)
	fmt.Println(org)
	return urlData.Original
}
