-- +goose Up
CREATE TABLE refresh_token (
  token TEXT PRIMARY KEY NOT NULL,
  user_id INTEGER,
  CONSTRAINT fk_refresh_token_user
    FOREIGN KEY (user_id) REFERENCES "user" (id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE refresh_token;
