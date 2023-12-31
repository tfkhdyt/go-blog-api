-- migrate:up
CREATE TABLE change_email_token (
  token TEXT PRIMARY KEY NOT NULL,
  new_email VARCHAR(255) NOT NULL,
  expires_at TIMESTAMP NOT NULL,
  user_id INTEGER,
  CONSTRAINT fk_change_email_token_user
    FOREIGN KEY (user_id) REFERENCES "user" (id)
    ON DELETE CASCADE
);

-- migrate:down
DROP TABLE change_email_token;
