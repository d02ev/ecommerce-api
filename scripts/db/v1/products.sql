CREATE TABLE IF NOT EXISTS products (
  id SERIAL CONSTRAINT pk_products_id PRIMARY KEY,
  created_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(255) NOT NULL,
  description TEXT NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  sku VARCHAR(50) NOT NULL,
  stock_qty INT NOT NULL,
  category_id INT NOT NULL,
  CONSTRAINT fk_products_category_id FOREIGN KEY (category_id) REFERENCES categories (id),
  CONSTRAINT uq_fk_products_category_id UNIQUE (category_id),
  CONSTRAINT uq_products_sku UNIQUE (sku)
);