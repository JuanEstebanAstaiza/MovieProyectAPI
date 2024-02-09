-- Crear trigger para actualizar el contador de visualizaciones en la tabla movies
DELIMITER //

CREATE TRIGGER after_insert_comment
    AFTER INSERT ON comments
    FOR EACH ROW
BEGIN
    UPDATE movies
    SET visualizations = visualizations + 1
    WHERE id = NEW.movie_id;
END;

//

DELIMITER ;