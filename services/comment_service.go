package services

import "github.com/JuanEstebanAstaiza/MovieProyectAPI/models"

func AddComment(comment models.Comment) error {
	// Agregar un comentario a la base de datos (esto es un ejemplo, debes adaptarlo según tu esquema)
	err := addCommentToDB(comment)
	if err != nil {
		return err
	}
	return nil
}

func DeleteComment(commentID int) error {
	// Eliminar un comentario de la base de datos (esto es un ejemplo, debes adaptarlo según tu esquema)
	err := deleteCommentFromDB(commentID)
	if err != nil {
		return err
	}
	return nil
}

func EditComment(updatedComment models.Comment) error {
	// Editar un comentario en la base de datos (esto es un ejemplo, debes adaptarlo según tu esquema)
	err := editCommentInDB(updatedComment)
	if err != nil {
		return err
	}
	return nil
}

// Funciones de ejemplo para interactuar con la base de datos
func addCommentToDB(comment models.Comment) error {
	// Implementar la lógica para agregar un comentario a la base de datos
	// Aquí puedes utilizar el objeto 'db' configurado en utils/db.go para interactuar con la base de datos
	_, err := db.Exec("INSERT INTO comments (user_id, content, create_time) VALUES (?, ?, NOW())",
		comment.UserID, comment.Content)
	if err != nil {
		return err
	}
	return nil
}

func deleteCommentFromDB(commentID int) error {
	// Implementar la lógica para eliminar un comentario de la base de datos
	// Aquí puedes utilizar el objeto 'db' configurado en utils/db.go para interactuar con la base de datos
	_, err := db.Exec("DELETE FROM comments WHERE id = ?", commentID)
	if err != nil {
		return err
	}
	return nil
}

func editCommentInDB(updatedComment models.Comment) error {
	// Implementar la lógica para editar un comentario en la base de datos
	// Aquí puedes utilizar el objeto 'db' configurado en utils/db.go para interactuar con la base de datos
	_, err := db.Exec("UPDATE comments SET content = ?, create_time = NOW() WHERE id = ?",
		updatedComment.Content, updatedComment.ID)
	if err != nil {
		return err
	}
	return nil
}
