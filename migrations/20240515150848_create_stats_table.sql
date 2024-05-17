-- +goose Up
-- +goose StatementBegin
CREATE TABLE teams (
    id SERIAL PRIMARY KEY,
    name VARCHAR(255) NOT NULL UNIQUE
);
CREATE TABLE players (
    id SERIAL PRIMARY KEY,
    first_name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    team_id INT NOT NULL,
    FOREIGN KEY (team_id) REFERENCES teams(id)
);
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
    player_id INT NOT NULL,
    FOREIGN KEY (player_id) REFERENCES players(id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE stats;
DROP TABLE players;
DROP TABLE teams;
-- +goose StatementEnd
