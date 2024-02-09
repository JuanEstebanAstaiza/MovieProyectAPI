-- Crear procedimiento almacenado para obtener los comentarios de una película
DELIMITER //

CREATE PROCEDURE get_movie_comments(IN movie_id INT)
BEGIN
    SELECT id AS comment_id, user_id, comment AS comment_text, timestamp
    FROM comments
    WHERE movie_id = movie_id;
END;

//

DELIMITER ;
