package models

type URLData struct {
	Id       int    `json:"id"`
	Original string `json:"org_url"`
	Shortend string `json:"short_url"`
}

// CREATE TABLE urldata (
//     id SERIAL PRIMARY KEY,
//     original TEXT NOT NULL,
//     shortend TEXT UNIQUE NOT NULL
// );
