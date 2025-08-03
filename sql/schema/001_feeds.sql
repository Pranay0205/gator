-- +goose Up
CREATE TABLE feeds(
  id UUID PRIMARY KEY DEFAULT gen_random_uuid(),
  created_at TIMESTAMP NOT NULL,
  updated_at TIMESTAMP NOT NULL,
  name VARCHAR(100) UNIQUE NOT NULL,
  url VARCHAR(255) UNIQUE NOT NULL,
  user_id UUID NOT NULL,
  CONSTRAINT fk_user
    FOREIGN KEY(user_id) 
    REFERENCES users(id)
    ON DELETE CASCADE
);

-- +goose Down
DROP TABLE feeds;