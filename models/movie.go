package models

// Movie estructura interna de la aplicación
type Movie struct {
	ID               *int    `json:"id"`
	Title            *string `json:"title"`
	Overview         *string `json:"overview"`
	ReleaseDate      *string `json:"release_date"`
	OriginalLanguage *string `json:"original_language"`
	// Otros campos según sea necesario
}

// ExternalMovie estructura para deserializar datos de la API externa (TMDb)
type ExternalMovie struct {
	ID               *int    `json:"id"`
	Title            *string `json:"title"`
	Overview         *string `json:"overview"`
	ReleaseDate      *string `json:"release_date"`
	OriginalLanguage *string `json:"original_language"`
	// Otros campos según sea necesario
}

// ConvertFromExternalMovie convierte una ExternalMovie a Movie
func ConvertFromExternalMovie(externalMovie ExternalMovie) Movie {
	return Movie{
		ID:               externalMovie.ID,
		Title:            externalMovie.Title,
		Overview:         externalMovie.Overview,
		ReleaseDate:      externalMovie.ReleaseDate,
		OriginalLanguage: externalMovie.OriginalLanguage,
		// Otros campos según sea necesario
	}
}
