-- +goose Up
-- +goose StatementBegin
ALTER TABLE sessions ADD COLUMN referral_link VARCHAR(255) UNIQUE;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE sessions DROP COLUMN referral_link;
-- +goose StatementEnd