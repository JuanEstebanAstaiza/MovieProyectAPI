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

// Funciones de ejemplo para interactuar con la base de datos
func saveMovieDetailsToDB(movieDetails models.Movie) error {
	// Implementar la lógica para guardar los detalles de la película en la base de datos
	// Aquí puedes utilizar el objeto 'db' configurado en utils/db.go para interactuar con la base de datos
	_, err := db.Exec("INSERT INTO movies (title, overview, release_date, original_language) VALUES (?, ?, ?, ?)",
		movieDetails.Title, movieDetails.Overview, movieDetails.ReleaseDate, movieDetails.OriginalLanguage)
	if err != nil {
		return err
	}
	return nil
}

func incrementMovieViewCountInDB(movieID string) error {
	// Implementar la lógica para incrementar el contador de visualizaciones en la base de datos
	// Aquí puedes utilizar el objeto 'db' configurado en utils/db.go para interactuar con la base de datos
	_, err := db.Exec("UPDATE movies SET view_count = view_count + 1 WHERE id = ?", movieID)
	if err != nil {
		return err
	}
	return nil
}

func getMostViewedMoviesFromDB(n int) ([]models.Movie, error) {
	// Implementar la lógica para obtener las n películas más visualizadas desde la base de datos
	// Aquí puedes utilizar el objeto 'db' configurado en utils/db.go para interactuar con la base de datos
	rows, err := db.Query("SELECT * FROM movies ORDER BY view_count DESC LIMIT ?", n)
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
