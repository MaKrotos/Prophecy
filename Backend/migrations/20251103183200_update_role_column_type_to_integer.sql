-- +goose Up
-- +goose StatementBegin
-- Обновляем существующие записи, преобразуя строковые значения ролей в числовые
UPDATE telegram_users SET role = '0' WHERE role = '' OR role IS NULL;
UPDATE telegram_users SET role = '1' WHERE role = 'Архитектор';
-- Изменяем тип столбца role на INTEGER с дефолтным значением 0
-- Сначала удаляем дефолтное значение, чтобы избежать ошибки преобразования
ALTER TABLE telegram_users ALTER COLUMN role DROP DEFAULT;
ALTER TABLE telegram_users ALTER COLUMN role TYPE INTEGER USING role::INTEGER;
ALTER TABLE telegram_users ALTER COLUMN role SET DEFAULT 0;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
-- Возвращаем тип столбца role к VARCHAR(50)
ALTER TABLE telegram_users ALTER COLUMN role DROP DEFAULT;
ALTER TABLE telegram_users ALTER COLUMN role TYPE VARCHAR(50) USING role::VARCHAR(50);
ALTER TABLE telegram_users ALTER COLUMN role SET DEFAULT '';
-- Восстанавливаем строковые значения ролей
UPDATE telegram_users SET role = '' WHERE role = 0;
UPDATE telegram_users SET role = 'Архитектор' WHERE role = 1;
-- +goose StatementEnd