CREATE TABLE IF NOT EXISTS categories (
  id SERIAL CONSTRAINT pk_categories_id PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(100) NOT NULL,
  description TEXT,
  CONSTRAINT uq_categories_name UNIQUE (name)
);
CREATE INDEX idx_categories_name ON categories (name);