package handler

import (
	"database/sql"
	"encoding/json"
	"github.com/andikanugraha11/Golang-Boilerplate-awesome-echo/app/repository"
	"github.com/andikanugraha11/Golang-Boilerplate-awesome-echo/app/repository/dev"
	"github.com/labstack/echo/v4"
	"net/http"
)

// NewPostHandler ...
func NewDevHandler(db *sql.DB) *Dev {
	return &Dev{
		repo: dev.NewSQLDevRepo(db),
	}
}

// Post ...
type Dev struct {
	repo repository.DevRepo
}

// Fetch all post data
func (p *Dev) Fetch(c echo.Context) {
	payload, _ := p.repo.Fetch(c, 5)
	response := utils.JsonResponse(true,"success", payload)
	return c.JSON(http.StatusOK, response)
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