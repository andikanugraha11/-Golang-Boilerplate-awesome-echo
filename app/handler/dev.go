package handler

import (
	"encoding/json"
	"net/http"
)

// NewPostHandler ...
func NewPostHandler(db *driver.DB) *Post {
	return &Post{
		repo: post.NewSQLPostRepo(db.SQL),
	}
}

// Post ...
type Post struct {
	repo repository.PostRepo
}

// Fetch all post data
func (p *Post) Fetch(w http.ResponseWriter, r *http.Request) {
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