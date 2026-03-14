CREATE TABLE roles
(
    id    BIGSERIAL PRIMARY KEY,
    name  VARCHAR(32) NOT NULL UNIQUE,
    title VARCHAR(64)
);

ALTER TABLE users
ADD COLUMN role_id BIGINT,
ADD CONSTRAINT fk_role_id
FOREIGN KEY (role_id) REFERENCES roles(id) ON DELETE SET NULL;

INSERT INTO roles (name, title) VALUES
                                    ('user', 'Пользователь'),
                                    ('admin', 'Администратор')
