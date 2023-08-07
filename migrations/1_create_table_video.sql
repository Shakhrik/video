CREATE TABLE IF NOT EXISTS videos(
    id VARCHAR(32) PRIMARY KEY,
    title VARCHAR(255),
    user_id VARCHAR(255),
    description text,
    year INT,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP
);