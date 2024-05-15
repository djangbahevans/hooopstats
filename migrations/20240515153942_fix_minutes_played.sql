-- +goose Up
-- +goose StatementBegin
ALTER TABLE stats
    ALTER COLUMN minutes_played TYPE REAL USING minutes_played::REAL;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stats
    ALTER COLUMN minutes_played TYPE INTEGER USING minutes_played::INTEGER;
-- +goose StatementEnd
