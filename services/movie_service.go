package services

import (
	"errors"
	"fmt"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/utils"
)

func MovieExists(ID string) (bool, error) {
	rows, err := utils.DB.Query("SELECT 1 FROM movies WHERE api_id = ?", ID)
	if err != nil {
		return false, err
	}
	defer rows.Close()

	if !rows.Next() {
		if err := rows.Err(); err != nil {
			return false, err
		}
		return false, nil
	}

	return true, nil
}

// GetMovieDetails obtiene los detalles de la película por su ID,
// guarda los detalles en la base de datos y luego incrementa el contador de visualizaciones.
func GetMovieDetails(movieID string) (models.Movie, error) {
	exists, err := MovieExists(movieID)
	if err != nil {
		return models.Movie{}, fmt.Errorf("error al verificar la base de datos: %v", err)
	}

	// Obtener detalles de la película desde la API de TMDb
	if !exists {
		movieDetails, err := utils.GetMovieDetailsFromTMDb(movieID)
		if err != nil {
			return models.Movie{}, fmt.Errorf("error al obtener detalles de la película: %v", err)
		}

		err = SaveMovieDetailsAndIncrementViewCount(movieDetails, movieID)
		if err != nil {
			return models.Movie{}, fmt.Errorf("error al guardar detalles de la película y aumentar el contador de visualizaciones: %v", err)
		}

		return movieDetails, nil
	}

	// Guardar los detalles de la película en la base de datos y aumentar el contador de visualizaciones
	err = IncrementViewCount(movieID)
	if err != nil {
		return models.Movie{}, fmt.Errorf("error al aumentar el contador de visualizaciones: %v", err)
	}

	rows, err := utils.DB.Query("SELECT api_id, title, overview, release_date, original_language FROM movies WHERE api_id = ?", movieID)
	if err != nil {
		return models.Movie{}, err
	} else {
		fmt.Println(err)
	}
	defer rows.Close()

	if !rows.Next() {
		return models.Movie{}, errors.New("person not found")
	}
	var movie models.Movie
	if err := rows.Scan(&movie.ID, &movie.Title, &movie.Overview, &movie.ReleaseDate, &movie.OriginalLanguage); err != nil {
		return models.Movie{}, err
	}
	return movie, nil
}

// IncrementViewCount incrementa el contador de visualizaciones de una película en la base de datos.
func IncrementViewCount(movieID string) error {
	// Incrementar el contador de visualizaciones en la base de datos
	_, err := utils.DB.Exec("UPDATE movies SET visualizations = visualizations + 1 WHERE api_id = ?", movieID)
	if err != nil {
		return fmt.Errorf("error al incrementar el contador de visualizaciones: %v", err)
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

// SaveMovieDetailsAndIncrementViewCount guarda los detalles de la película en la base de datos
// y aumenta el contador de visualizaciones. Retorna el ID único de la película.
func SaveMovieDetailsAndIncrementViewCount(movieDetails models.Movie, apiID string) error {
	// Generar un ID único para la película
	movieID, err := utils.GenerateUniqueMovieID()
	if err != nil {
		return fmt.Errorf("error al generar un ID único para la película: %v", err)
	}

	// Guardar los detalles de la película en la base de datos
	_, err = utils.DB.Exec("INSERT INTO movies (id, title, overview, release_date, original_language, api_id) VALUES (?, ?, ?, ?, ?, ?)",
		movieID, movieDetails.Title, movieDetails.Overview, movieDetails.ReleaseDate, movieDetails.OriginalLanguage, apiID)
	if err != nil {
		return fmt.Errorf("error al guardar los detalles de la película en la base de datos: %v", err)
	}

	// Incrementar el contador de visualizaciones en la base de datos
	err = IncrementViewCount(movieID)
	if err != nil {
		return fmt.Errorf("error al incrementar el contador de visualizaciones: %v", err)
	}

	return nil
}

func getMostViewedMoviesFromDB(n int) ([]models.Movie, error) {
	rows, err := utils.DB.Query("SELECT api_id, title, overview, release_date, original_language FROM movies ORDER BY visualizations DESC LIMIT ?", n)
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
