
CREATE TABLE users (
    id SERIAL PRIMARY KEY,
    username VARCHAR(64) NOT NULL,
    email VARCHAR(128) UNIQUE NOT NULL,
    password_digest VARCHAR(128) NOT NULL,
    is_admin BOOLEAN NOT NULL DEFAULT 'false'
);


