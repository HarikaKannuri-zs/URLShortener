package models

type URLData struct {
	Id       int    `json:"id"`
	Original string `json:"org_url"`
	Shortend string `json:"short_url"`
}

// type URLShortenReq struct {
// 	Original string `json:"org_url"`
// 	Alias    string `json:"alias"`
// }
