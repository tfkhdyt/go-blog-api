CREATE TABLE refresh_token (
  token TEXT PRIMARY KEY NOT NULL,
  user_id INTEGER,
  CONSTRAINT fk_refresh_token_user
    FOREIGN KEY (user_id) REFERENCES users (id)
    ON DELETE CASCADE
);
