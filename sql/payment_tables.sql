-- Tabla para almacenar información sobre los pagos
CREATE TABLE payments (
                          id INT AUTO_INCREMENT PRIMARY KEY,
                          user_id varchar(255),
                          amount DECIMAL(10, 2) NOT NULL,
                          description VARCHAR(255),
                          status ENUM('pending', 'success', 'failed') DEFAULT 'pending',
                          FOREIGN KEY (user_id) REFERENCES users(id)
);
