package store

import (
	"database/sql"
	"fmt"
	"url-shortener/cache"
	"url-shortener/models"
	_ "github.com/lib/pq"
)

type Store struct {
	db    *sql.DB
	Cache *cache.RedisCache
}

const DefaultDomain = "https://mygourl/"

func NewStore(urlDb *sql.DB, redisCache *cache.RedisCache) *Store {
	return &Store{
		db:    urlDb,
		Cache: redisCache,
	}
}

func (s *Store) ShortenUrl(urlData *models.URLData) (string, error) {
	fmt.Println("Saving alias:", urlData.Alias)
	query := `INSERT into urldata(original,alias,expiry_at,max_click_limit) VALUES($1,$2,$3,$4)`
	_, err := s.db.Exec(query, urlData.Original, urlData.Alias, urlData.ExpiryAt, urlData.MaxClickLimit)
	if err != nil {
		fmt.Printf("Error inserting into DB: %v\n", err)
		return "", err
	}
	return DefaultDomain + urlData.Alias, nil
}

func (s *Store) RedirectUrl(aliasUrl string) (*models.URLData, error) {

	var urlData models.URLData
	query := `SELECT original,expiry_at,click_cnt,max_click_limit from urlData where Alias = $1`
	err := s.db.QueryRow(query, aliasUrl).Scan(&urlData.Original, &urlData.ExpiryAt, &urlData.ClickCnt, &urlData.MaxClickLimit)
	if err != nil {
		return nil, err
	}

	return &urlData, nil
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

func (s *Store) IncrementClickCount(alias string) error {
	_, err := s.db.Exec(`UPDATE urlData SET click_cnt = click_cnt+1 where alias = $1`, alias)
	return err
}
