-- +goose Up
-- +goose StatementBegin
-- Обновляем существующие записи, устанавливая пустую строку вместо NULL
UPDATE telegram_users SET role = '' WHERE role IS NULL;
-- Изменяем тип столбца role на VARCHAR(50) с дефолтным значением пустой строки
ALTER TABLE telegram_users ALTER COLUMN role TYPE VARCHAR(50);
ALTER TABLE telegram_users ALTER COLUMN role SET DEFAULT '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE telegram_users ALTER COLUMN role DROP DEFAULT;
-- +goose StatementEnd