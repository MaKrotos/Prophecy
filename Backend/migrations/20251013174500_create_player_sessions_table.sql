-- +goose Up
-- +goose StatementBegin
CREATE TABLE player_sessions (
    player_id INTEGER NOT NULL REFERENCES telegram_users(id) ON DELETE CASCADE,
    session_id INTEGER NOT NULL REFERENCES sessions(id) ON DELETE CASCADE,
    joined_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    PRIMARY KEY (player_id, session_id)
);

CREATE INDEX idx_player_sessions_player_id ON player_sessions(player_id);
CREATE INDEX idx_player_sessions_session_id ON player_sessions(session_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE player_sessions;
-- +goose StatementEnd