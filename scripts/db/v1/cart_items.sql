CREATE TABLE IF NOT EXISTS cart_items (
  id SERIAL CONSTRAINT pk_cart_items_id PRIMARY KEY,
  created_at TIMESTAMPTZ NOT NULL DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMPTZ  NOT NULL DEFAULT CURRENT_TIMESTAMP,
  cart_id INT NOT NULL,
  product_id INT NOT NULL,
  quantity INT NOT NULL,
  total_price DECIMAL(10, 2) NOT NULL,
  CONSTRAINT fk_cart_items_cart_id FOREIGN KEY (cart_id) REFERENCES carts (id) ON DELETE CASCADE,
  CONSTRAINT fk_cart_items_product_id FOREIGN KEY (product_id) REFERENCES products (id)
);
CREATE INDEX idx_cart_items_cart_id ON cart_items (cart_id);
CREATE INDEX idx_cart_items_product_id ON cart_items (product_id);