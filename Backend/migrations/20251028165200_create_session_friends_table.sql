-- +goose Up
-- +goose StatementBegin
CREATE TABLE session_friends (
    id SERIAL PRIMARY KEY,
    session_id INTEGER NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    user_id INTEGER NOT NULL REFERENCES telegram_users(id) ON DELETE CASCADE,
    friend_id INTEGER NOT NULL REFERENCES telegram_users(id) ON DELETE CASCADE,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    UNIQUE(session_id, user_id, friend_id)
);

CREATE INDEX idx_session_friends_session_id ON session_friends(session_id);
CREATE INDEX idx_session_friends_user_id ON session_friends(user_id);
CREATE INDEX idx_session_friends_friend_id ON session_friends(friend_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE session_friends;
-- +goose StatementEnd