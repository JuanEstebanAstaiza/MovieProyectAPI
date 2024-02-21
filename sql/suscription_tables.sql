CREATE TABLE IF NOT EXISTS subscriptions (
                                             id INT AUTO_INCREMENT PRIMARY KEY,
                                             user_id varchar(255),
                                             plan VARCHAR(255) NOT NULL,
                                             price INT NOT NULL,
                                             description TEXT,
                                             status ENUM('active', 'inactive') NOT NULL DEFAULT 'inactive',
                                             FOREIGN KEY (user_id) REFERENCES users(id)
);
