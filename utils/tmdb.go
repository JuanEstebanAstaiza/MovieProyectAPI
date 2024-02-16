package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"

	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
)

const (
	tmdbBaseURL = "https://api.themoviedb.org/3"
	apiKey      = "94d94ff8c72fad39bf863279c1275845"
)

// GetMovieDetailsFromTMDb obtiene detalles de una película desde TMDb
func GetMovieDetailsFromTMDb(movieID string) (models.Movie, error) {
	endpoint := fmt.Sprintf("%s/movie/%s", tmdbBaseURL, movieID)
	queryParams := url.Values{}
	queryParams.Add("api_key", apiKey)

	resp, err := http.Get(endpoint + "?" + queryParams.Encode())
	if err != nil {
		return models.Movie{}, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return models.Movie{}, fmt.Errorf("Error al obtener detalles de la película. Código de estado: %d", resp.StatusCode)
	}

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return models.Movie{}, err
	}

	var externalMovie models.ExternalMovie
	err = json.Unmarshal(body, &externalMovie)
	if err != nil {
		return models.Movie{}, err
	}

	// Convertir a la estructura interna de models
	internalMovie := models.ConvertFromExternalMovie(externalMovie)

	return internalMovie, nil
}
