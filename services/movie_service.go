package services

import (
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/utils"
)

func GetMovieDetails(movieID string) (models.Movie, error) {
	// Obtener detalles de la película desde la API de TMDb
	movieDetails, err := utils.GetMovieDetailsFromTMDb(movieID)
	if err != nil {
		return models.Movie{}, err
	}

	// Incrementar el contador de visualizaciones (esto es un ejemplo, debes adaptarlo según tu esquema)
	err = incrementMovieViewCountInDB(movieID)
	if err != nil {
		return models.Movie{}, err
	}

	// Guardar los detalles de la película en la base de datos (esto es un ejemplo, debes adaptarlo según tu esquema)
	err = saveMovieDetailsToDB(movieDetails)
	if err != nil {
		return models.Movie{}, err
	}

	return movieDetails, nil
}

func IncrementViewCount(movieID string) error {
	// Incrementar el contador de visualizaciones en la base de datos (esto es un ejemplo, debes adaptarlo según tu esquema)
	err := incrementMovieViewCountInDB(movieID)
	if err != nil {
		return err
	}
	return nil
}

func GetMostViewedMovies(n int) ([]models.Movie, error) {
	// Obtener las n películas más visualizadas desde la base de datos (esto es un ejemplo, debes adaptarlo según tu esquema)
	mostViewedMovies, err := getMostViewedMoviesFromDB(n)
	if err != nil {
		return nil, err
	}
	return mostViewedMovies, nil
}

func saveMovieDetailsToDB(movieDetails models.Movie) error {

	_, err := utils.DB.Exec("INSERT INTO movies (title, overview, release_date, original_language) VALUES (?, ?, ?, ?)",
		movieDetails.Title, movieDetails.Overview, movieDetails.ReleaseDate, movieDetails.OriginalLanguage)
	if err != nil {
		return err
	}
	return nil
}

func incrementMovieViewCountInDB(movieID string) error {

	_, err := utils.DB.Exec("UPDATE movies SET visualizations = visualizations + 1 WHERE id = ?", movieID)
	if err != nil {
		return err
	}
	return nil
}

func getMostViewedMoviesFromDB(n int) ([]models.Movie, error) {
	rows, err := utils.DB.Query("SELECT * FROM movies ORDER BY visualizations DESC LIMIT ?", n)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var mostViewedMovies []models.Movie
	for rows.Next() {
		var movie models.Movie
		err := rows.Scan(&movie.ID, &movie.Title, &movie.Overview, &movie.ReleaseDate, &movie.OriginalLanguage)
		if err != nil {
			return nil, err
		}
		mostViewedMovies = append(mostViewedMovies, movie)
	}

	return mostViewedMovies, nil
}
