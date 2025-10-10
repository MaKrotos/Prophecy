-- +goose Up
-- +goose StatementBegin
ALTER TABLE telegram_users ADD COLUMN generated_name VARCHAR(255);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE telegram_users DROP COLUMN IF EXISTS generated_name;
-- +goose StatementEnd