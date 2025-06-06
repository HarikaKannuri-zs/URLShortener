package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strings"
	"url-shortener/models"
	"url-shortener/service"
	"url-shortener/store"
)

type Handler struct {
	srvc *service.Service
}

func NewHandler(s *service.Service) *Handler {
	return &Handler{
		srvc: s,
	}
}

func (h *Handler) ShortenUrl(w http.ResponseWriter, r *http.Request) {
	var req *models.URLData
	err := json.NewDecoder(r.Body).Decode(&req)
	if err != nil {
		http.Error(w, "Decoding Error", http.StatusBadRequest)
		return
	}
	shortUrl, err := h.srvc.ShortenUrl(req)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(shortUrl)
	if err != nil {
		http.Error(w, "Failed to encode", http.StatusInternalServerError)
		return
	}

}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {

	reqUrl := r.URL.Query().Get("url")
	if reqUrl == "" {
		fmt.Println("Empty URL received")
		http.Error(w, "Invalid short URL format", http.StatusBadRequest)
		return
	}
	alias := strings.TrimPrefix(reqUrl, store.DefaultDomain)
	if alias == reqUrl {
		fmt.Println("URL is wrong")
		http.Error(w, "Invalid short URL format", http.StatusBadRequest)
		return
	}
	orgUrl := h.srvc.RedirectUrl(alias)
	if orgUrl == "" {
		fmt.Println("No original URL found...")
		http.Error(w, "Invalid Request", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, orgUrl, http.StatusFound)

}
