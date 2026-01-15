CREATE TABLE users (
   id BIGSERIAL PRIMARY KEY,
   email TEXT NOT NULL UNIQUE,
   password_hash TEXT NOT NULL,
   created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

CREATE UNIQUE INDEX idx_users_email ON users(email);
