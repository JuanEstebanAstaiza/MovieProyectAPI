-- Crear procedimiento almacenado para obtener los comentarios de una película
CREATE OR REPLACE FUNCTION get_movie_comments(movie_id INT)
    RETURNS TABLE (
                      comment_id INT,
                      user_id INT,
                      comment_text TEXT,
                      timestamp TIMESTAMP
                  ) AS $$
BEGIN
    RETURN QUERY
        SELECT id, user_id, comment_text, timestamp
        FROM comments
        WHERE movie_id = movie_id;
END;
$$ LANGUAGE plpgsql;
