-- +goose Up
-- +goose StatementBegin
ALTER TABLE telegram_users ADD COLUMN role VARCHAR(50) DEFAULT NULL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE telegram_users DROP COLUMN IF EXISTS role;
-- +goose StatementEnd