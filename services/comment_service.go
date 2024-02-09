package services

import (
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/models"
	"github.com/JuanEstebanAstaiza/MovieProyectAPI/utils"
)

func AddCommentToDB(comment models.Comment) error {
	_, err := utils.DB.Exec("INSERT INTO comments (user_id, comment_text, timestamp) VALUES (?, ?, NOW())",
		comment.UserID, comment.Content)
	if err != nil {
		return err
	}
	return nil
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
