CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL
);

-- Insert a sample user
INSERT INTO users (name, email) VALUES ('Alice In Chains', 'alice@example.com');
