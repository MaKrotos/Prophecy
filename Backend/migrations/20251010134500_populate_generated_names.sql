-- +goose Up
-- +goose StatementBegin
UPDATE telegram_users 
SET generated_name = 'Халяльная канарейка' 
WHERE generated_name IS NULL OR generated_name = '';
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
UPDATE telegram_users 
SET generated_name = NULL 
WHERE generated_name = 'Халяльная канарейка';
-- +goose StatementEnd