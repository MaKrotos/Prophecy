-- +goose Up
-- +goose StatementBegin
ALTER TABLE sessions ADD COLUMN expires_at TIMESTAMP;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sessions DROP COLUMN expires_at;
-- +goose StatementEnd