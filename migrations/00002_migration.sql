-- +goose Up
INSERT INTO "user" (id, email) VALUES (1, 'test@ozon.ru');

-- +goose Down
DELETE FROM "user" WHERE id = 1;
