CREATE TABLE IF NOT EXISTS users (
  id SERIAL CONSTRAINT pk_users_id PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(150) NOT NULL,
  role INT NOT NULL DEFAULT 0,
  password_hash VARCHAR(255) NOT NULL,
  refresh_token VARCHAR(255),
  CONSTRAINT uq_users_refresh_token UNIQUE (refresh_token),
  CONSTRAINT uq_users_email UNIQUE (email)
);