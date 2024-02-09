package models

// Movie estructura interna de la aplicación
type Movie struct {
	ID               *int    `json:"id"`
	Title            *string `json:"title"`
	Overview         *string `json:"overview"`
	ReleaseDate      *string `json:"release_date"`
	OriginalLanguage *string `json:"original_language"`
}

// ExternalMovie estructura para deserializar datos de la API externa (TMDb)
type ExternalMovie struct {
	ID               *int    `json:"id"`
	Title            *string `json:"title"`
	Overview         *string `json:"overview"`
	ReleaseDate      *string `json:"release_date"`
	OriginalLanguage *string `json:"original_language"`
}

// ConvertFromExternalMovie convierte una ExternalMovie a Movie
func ConvertFromExternalMovie(externalMovie ExternalMovie) Movie {
	return Movie{
		ID:               externalMovie.ID,
		Title:            externalMovie.Title,
		Overview:         externalMovie.Overview,
		ReleaseDate:      externalMovie.ReleaseDate,
		OriginalLanguage: externalMovie.OriginalLanguage,
	}
}
