CREATE TABLE tasks (
    id SERIAL PRIMARY KEY,
    name VARCHAR(100) NOT NULL,
    creator_id INT NOT NULL,
    created_at TIMESTAMP DEFAULT now() NOT NULL,
    updated_at TIMESTAMP DEFAULT now() NOT NULL,
    date_of_remind TIMESTAMP,
    date_of_appoint TIMESTAMP,
    date_of_complete TIMESTAMP,
    description TEXT,
    FOREIGN KEY (creator_id) REFERENCES users(id) ON DELETE CASCADE
)