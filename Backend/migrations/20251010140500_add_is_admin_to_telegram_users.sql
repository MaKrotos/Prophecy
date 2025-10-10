-- +goose Up
-- +goose StatementBegin
ALTER TABLE telegram_users ADD COLUMN is_admin BOOLEAN DEFAULT FALSE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE telegram_users DROP COLUMN IF EXISTS is_admin;
-- +goose StatementEnd