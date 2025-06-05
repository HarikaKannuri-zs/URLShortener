package handler

import (
	"encoding/json"
	"fmt"
	"net/http"
	"url-shortener/models"
	"url-shortener/service"
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
	err = h.srvc.ShortenUrl(req)
	if err != nil {
		http.Error(w, "Invalid Request", http.StatusBadRequest)
		return
	}
	w.WriteHeader(http.StatusCreated)

}

func (h *Handler) Redirect(w http.ResponseWriter, r *http.Request) {

	reqUrl := r.URL.Query().Get("url")
	if reqUrl == "" {
		fmt.Println("Empty URL received")
		http.Error(w, "Invalid Request", http.StatusInternalServerError)
		return
	}
	orgUrl := h.srvc.RedirectUrl(reqUrl)
	if orgUrl == "" {
		fmt.Println("No original URL found...")
		http.Error(w, "Invalid Request", http.StatusNotFound)
		return
	}
	http.Redirect(w, r, orgUrl, http.StatusFound)

}
