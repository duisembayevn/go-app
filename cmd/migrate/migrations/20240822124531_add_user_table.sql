-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `users`
(
    `id`        INT          NOT NULL AUTO_INCREMENT,
    `firstName` VARCHAR(50)  NOT NULL,
    `lastName`  VARCHAR(50)  NOT NULL,
    `username`  VARCHAR(50)  NOT NULL UNIQUE,
    `password`  VARCHAR(255) NOT NULL,
    PRIMARY KEY (`id`)
);


-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table user;
-- +goose StatementEnd
