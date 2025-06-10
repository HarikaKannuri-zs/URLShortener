package models

import "time"

type URLData struct {
	Id            int       `json:"id"`
	Original      string    `json:"org_url"`
	Alias         string    `json:"alias,omitempty"`
	CretedAt      time.Time `json:"-"`
	ExpiryAt      time.Time `json:"expiry_at"`
	ClickCnt      int       `json:"-"`
	MaxClickLimit int       `json:"max_click_limit"`
}

type URLDataCache struct{
	
}

// CREATE TABLE urlData(
// 	id SERIAL  PRIMARY KEY,
// 	original TEXT NOT NULL,
// 	alias TEXT UNIQUE NOT NULL,
// 	created_at TIMESTAMP DEFAULT NOW(),
// 	expiry_at TIMESTAMP,
// 	click_cnt INT DEFAULT 0,
// 	max_click_limit INT
// );
