ALTER TABLE users ADD COLUMN full_name TEXT;

UPDATE users SET full_name = name;
