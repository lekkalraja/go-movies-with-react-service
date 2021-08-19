package main

import (
	"net/http"
	"strconv"
	"time"

	"github.com/julienschmidt/httprouter"

	"github.com/lekkalraja/go-movies-with-react-service/models"
)

func (app *application) getMovie(w http.ResponseWriter, r *http.Request) {
	params := httprouter.ParamsFromContext(r.Context())

	id, err := strconv.Atoi(params.ByName("id"))

	if err != nil {
		data := models.AppError{
			ErrCode: http.StatusBadRequest,
			ErrDesc: err.Error(),
		}
		writeJSON(w, http.StatusBadRequest, data, "error")
		return
	}
	movie := getMovie(id)
	writeJSON(w, http.StatusOK, movie, "movie")
}

func (app *application) getMovies(w http.ResponseWriter, r *http.Request) {
	var movies []models.Movie
	movies = append(movies, getMovie(12))
	movies = append(movies, getMovie(13))
	movies = append(movies, getMovie(14))
	writeJSON(w, http.StatusOK, movies, "movies")
}

func getMovie(id int) models.Movie {
	return models.Movie{
		ID:          id,
		Title:       "GangLeader",
		Description: "Biggest Blockbuster for Chiranjeevi",
		ReleaseDate: time.Date(1993, 12, 4, 12, 9, 12, 0, time.UTC),
		Runtime:     135,
		Rating:      10,
		MPAARating:  "EX",
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}
}
