CREATE TABLE IF NOT EXISTS roles (
  id SERIAL CONSTRAINT pk_roles_id PRIMARY KEY,
  created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
  name VARCHAR(100) NOT NULL,
  description TEXT,
  user_id INT NOT NULL,
  CONSTRAINT uq_roles_name UNIQUE (name),
  CONSTRAINT fk_roles_user_id FOREIGN KEY (user_id) REFERENCES users (id)
);
CREATE INDEX idx_roles_user_id ON roles (user_id);
CREATE INDEX idx_roles_name ON roles (name);