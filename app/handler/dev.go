package handler

import (
	"encoding/json"
	"github.com/andikanugraha11/Golang-Boilerplate-awesome-echo/app/config"
	"github.com/andikanugraha11/Golang-Boilerplate-awesome-echo/app/repository"
	"github.com/andikanugraha11/Golang-Boilerplate-awesome-echo/app/repository/dev"
	"net/http"
)

// NewPostHandler ...
func NewPostHandler(db *config.DB) *Dev {
	return &Dev{
		repo: dev.NewSQLDevRepo(db.SQL),
	}
}

// Post ...
type Dev struct {
	repo repository.DevRepo
}

// Fetch all post data
func (p *Dev) Fetch(w http.ResponseWriter, r *http.Request) {
	payload, _ := p.repo.Fetch(r.Context(), 5)

	respondWithJSON(w, http.StatusOK, payload)
}

func respondWithJSON(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJSON(w, code, map[string]string{"message": msg})
}