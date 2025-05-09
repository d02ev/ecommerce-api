CREATE TABLE IF NOT EXISTS addresses (
  id SERIAL CONSTRAINT pk_addresses_id PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  user_id INT NOT NULL,
  add_line_1 VARCHAR(255) NOT NULL,
  add_line_2 VARCHAR(255),
  landmark VARCHAR(255),
  city VARCHAR(100) NOT NULL,
  state VARCHAR(100) NOT NULL,
  zip_code VARCHAR(20) NOT NULL,
  country VARCHAR(100) NOT NULL,
  CONSTRAINT fk_addresses_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);
CREATE INDEX idx_addresses_user_id ON addresses(user_id);