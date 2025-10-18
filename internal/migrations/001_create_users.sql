CREATE TABLE IF NOT EXISTS users (
    id INTEGER PRIMARY KEY AUTOINCREMENT,
    name TEXT NOT NULL,
    email TEXT UNIQUE NOT NULL,
    full_name TEXT
);

-- Insert a sample user
INSERT OR IGNORE INTO users (name, email) VALUES ('Alice In Chains', 'alice@example.com');
