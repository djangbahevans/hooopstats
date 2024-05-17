-- +goose Up
-- +goose StatementBegin
ALTER TABLE stats ADD COLUMN game_number INT NOT NULL;
ALTER TABLE stats ADD CONSTRAINT unique_game_number_per_player UNIQUE (player_id, game_number);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
ALTER TABLE stats DROP CONSTRAINT unique_game_number_per_player;
-- +goose StatementEnd
