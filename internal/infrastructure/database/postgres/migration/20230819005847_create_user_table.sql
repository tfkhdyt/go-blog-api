-- migrate:up
CREATE TYPE role AS ENUM ('user', 'admin');
CREATE TABLE "user" (
  id SERIAL PRIMARY KEY NOT NULL,
  full_name VARCHAR(50) NOT NULL,
  username VARCHAR(16) NOT NULL UNIQUE,
  email VARCHAR(255) NOT NULL UNIQUE,
  password TEXT NOT NULL,
  role role DEFAULT 'user',
  created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP
);

-- migrate:down
DROP TABLE "user";
DROP TYPE role;
