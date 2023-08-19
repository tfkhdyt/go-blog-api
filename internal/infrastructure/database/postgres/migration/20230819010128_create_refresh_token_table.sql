-- migrate:up
CREATE TABLE refresh_token (
  token TEXT PRIMARY KEY NOT NULL,
  user_id INTEGER,
  CONSTRAINT fk_refresh_token_user
    FOREIGN KEY (user_id) REFERENCES "user" (id)
    ON DELETE CASCADE
);

-- migrate:down
DROP TABLE refresh_token;
