-- +goose Up
-- SQL in section 'Up' is executed when this migration is applied
CREATE TABLE history
(
    id         SERIAL PRIMARY KEY,
    user_id    INT          NOT NULL,
    action     VARCHAR(50)  NOT NULL,
    item       VARCHAR(100) NOT NULL,
    quantity   INT          NOT NULL,
    timestamp  TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP    NOT NULL DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY (user_id) REFERENCES users (id)
);

-- +goose Down
-- SQL section 'Down' is executed when this migration is rolled back
DROP TABLE IF EXISTS history;
