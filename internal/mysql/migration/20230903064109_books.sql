-- migrate:up
    CREATE TABLE books(
        id INT NOT NULL AUTO_INCREMENT PRIMARY KEY,
        title VARCHAR(100) NOT NULL,
        author VARCHAR(255) NOT NULL DEFAULT "anonymous"
    )

-- migrate:down
DROP TABLE IF EXISTS books;

