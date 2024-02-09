-- Crear la tabla de usuarios
CREATE TABLE users (
                       id INT PRIMARY KEY,
                       nickname VARCHAR(50) NOT NULL,
                       email VARCHAR(100) NOT NULL,
                       password VARCHAR(255) NOT NULL


);

-- Crear la tabla de películas
CREATE TABLE movies (
                        id INT PRIMARY KEY,
                        api_id INT NOT NULL,
                        visualizations INT DEFAULT 0, -- default value = 0
                        title VARCHAR(255) NOT NULL,
                        overview TEXT,
                        release_date DATE,
                        original_language VARCHAR(10)

);

-- Crear la tabla de comentarios
CREATE TABLE comments (
                          id INT PRIMARY KEY,
                          user_id INT,
                          movie_id INT,
                          comment_text TEXT,
                          timestamp TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
                          FOREIGN KEY (user_id) REFERENCES users(id),
                          FOREIGN KEY (movie_id) REFERENCES movies(id)

);
