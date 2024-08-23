-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `projects` (
    `id` INT NOT NULL AUTO_INCREMENT,
    `name` VARCHAR(100) NOT NULL,
    `desc` TEXT,
    `userId` INT NOT NULL,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`userId`) REFERENCES `users`(`id`) ON DELETE CASCADE
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE projects
-- +goose StatementEnd
