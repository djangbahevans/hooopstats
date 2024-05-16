-- +goose Up
-- +goose StatementBegin
CREATE TABLE stats (
    id SERIAL PRIMARY KEY,
    points INT NOT NULL,
    rebounds INT NOT NULL,
    assists INT NOT NULL,
    steals INT NOT NULL,
    blocks INT NOT NULL,
    fouls INT NOT NULL,
    turnovers INT NOT NULL,
    minutes_played INT NOT NULL,
    player_id INT NOT NULL
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stats;
-- +goose StatementEnd
