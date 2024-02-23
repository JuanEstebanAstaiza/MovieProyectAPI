package services

import (
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/utils"
	"github.com/google/uuid"
)

// AddCommentToDB inserta un comentario en la base de datos.
func AddCommentToDB(comment models.Comment) error {
	// Verificar si la película con el ID proporcionado existe en la base de datos
	var movieID int
	err := utils.DB.QueryRow("SELECT id FROM movies WHERE id = ?", comment.MovieID).Scan(&movieID)
	if err != nil {
		return err // La película con el ID proporcionado no existe
	}

	// Generar un ID único de comentario
	commentUUID := uuid.New()
	commentID := HashUUID(commentUUID.String())

	// Insertar el comentario en la base de datos con el ID único generado
	_, err = utils.DB.Exec("INSERT INTO comments (id, movie_id, user_id, comment_text, timestamp) VALUES (?, ?, ?, ?, NOW())",
		commentID, comment.MovieID, comment.UserID, comment.Content)
	if err != nil {
		return err
	}
	return nil
}

// HashUUID convierte un UUID como cadena en un valor hash de tipo INT.
func HashUUID(uuidStr string) int {
	hash := 0
	for _, char := range uuidStr {
		hash = 31*hash + int(char)
	}
	return hash
}

func DeleteCommentFromDB(commentID int) error {
	_, err := utils.DB.Exec("DELETE FROM comments WHERE id = ?", commentID)
	if err != nil {
		return err
	}
	return nil
}

func EditCommentInDB(updatedComment models.Comment) error {

	_, err := utils.DB.Exec("UPDATE comments SET comment_text = ?, timestamp = NOW() WHERE id = ?",
		updatedComment.Content, updatedComment.ID)
	if err != nil {
		return err
	}
	return nil
}
