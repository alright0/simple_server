ALTER TABLE users
DROP COLUMN is_deleted,
DROP COLUMN updated_at;

-- Удаление триггера обновления updated_at для users
DROP TRIGGER IF EXISTS trg_users_updated_at ON users;

-- Удаление функции обновления столбца updated_at
DROP FUNCTION IF EXISTS set_updated_at();