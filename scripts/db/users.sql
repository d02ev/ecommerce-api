CREATE TABLE IF NOT EXISTS users (
  id SERIAL CONSTRAINT pk_users_id PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(100) NOT NULL,
  email VARCHAR(150) NOT NULL,
  password_hash VARCHAR(60) NOT NULL,
  refresh_token VARCHAR(255) NOT NULL,
  CONSTRAINT uq_users_password_hash UNIQUE (password_hash),
  CONSTRAINT uq_users_refresh_token UNIQUE (refresh_token),
  CONSTRAINT uq_users_email UNIQUE (email)
);
CREATE INDEX idx_users_email ON users (email);