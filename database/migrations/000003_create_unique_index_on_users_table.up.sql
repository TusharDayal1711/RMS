CREATE UNIQUE INDEX IF NOT EXISTS unique_users_name_email_id
ON users(name, email, id);
