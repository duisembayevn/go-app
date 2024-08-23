-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS `tasks` (
    `id` Int not null auto_increment,
    `name` varchar(255) not null,
    `completed` Boolean not null default false,
    `projectId` int not null,

    PRIMARY KEY (`id`),
    FOREIGN KEY (`projectId`) REFERENCES `projects`(`id`) ON DELETE CASCADE
);

-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
drop table tasks
-- +goose StatementEnd
