CREATE TABLE IF NOT EXISTS products (
  id SERIAL CONSTRAINT pk_products_id PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(100) NOT NULL,
  description TEXT NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  sku VARCHAR(50) NOT NULL UNIQUE,
  stock_qty INT NOT NULL,
  category_id INT NOT NULL,
  CONSTRAINT fk_products_category_id FOREIGN KEY (category_id) REFERENCES categories (id)
);
CREATE INDEX idx_products_category_id ON products (category_id);
CREATE INDEX idx_products_sku ON products (sku);