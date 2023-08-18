CREATE TABLE reset_password_token (
  token TEXT PRIMARY KEY NOT NULL,
  expires_at TIMESTAMP NOT NULL,
  user_id INTEGER,
  CONSTRAINT fk_reset_password_token_user
    FOREIGN KEY (user_id) REFERENCES users (id)
    ON DELETE CASCADE
);

