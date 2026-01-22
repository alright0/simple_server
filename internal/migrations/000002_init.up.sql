ALTER TABLE users
ADD COLUMN is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
ADD COLUMN updated_at TIMESTAMPTZ;

-- создание функции автообновления при обновлении столбца updated_at
CREATE OR REPLACE FUNCTION set_updated_at()
RETURNS TRIGGER AS $$
BEGIN
  IF NEW IS DISTINCT FROM OLD THEN
    NEW.updated_at = now();
END IF;
RETURN NEW;
END;
$$ LANGUAGE plpgsql;

-- Создание триггера обноваления столбца updated_at для users
DO $$
BEGIN
  IF NOT EXISTS (
    SELECT 1
    FROM pg_trigger
    WHERE tgname = 'trg_users_updated_at'
  ) THEN
CREATE TRIGGER trg_users_updated_at
    BEFORE UPDATE ON users
    FOR EACH ROW
    EXECUTE FUNCTION set_updated_at();
END IF;
END;
$$;