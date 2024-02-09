-- Crear trigger para actualizar el contador de visualizaciones en la tabla movies
CREATE OR REPLACE FUNCTION update_visualizations()
    RETURNS TRIGGER AS $$
BEGIN
    UPDATE movies
    SET visualizations = visualizations + 1
    WHERE id = NEW.movie_id;
    RETURN NEW;
END;
$$ LANGUAGE plpgsql;

CREATE TRIGGER after_insert_comment
    AFTER INSERT ON comments
    FOR EACH ROW
EXECUTE FUNCTION update_visualizations();
