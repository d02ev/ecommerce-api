CREATE TABLE IF NOT EXISTS orders (
  id SERIAL CONSTRAINT pk_orders_id PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  user_id INT NOT NULL,
  status VARCHAR(50) NOT NULL,
  total_amt DECIMAL(10, 2) NOT NULL,
  shipping_address_id INT NOT NULL,
  CONSTRAINT fk_orders_user_id FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE,
  CONSTRAINT fk_orders_shipping_address_id FOREIGN KEY (shipping_address_id) REFERENCES addresses (id)
);
CREATE INDEX idx_orders_user_id ON orders (user_id);
CREATE INDEX idx_orders_shipping_address_id ON orders (shipping_address_id);