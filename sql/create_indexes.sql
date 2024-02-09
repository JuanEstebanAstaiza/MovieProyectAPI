-- Crear índice en el campo email de la tabla users (para mejorar búsquedas por email)
CREATE INDEX idx_users_email ON users(email);

-- Otros índices según sea necesario
