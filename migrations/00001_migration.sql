-- +goose Up
CREATE TABLE "user"(
    id SERIAL,
    email varchar(255)
);

ALTER TABLE "user" ADD CONSTRAINT "user_primary_key_id" PRIMARY KEY (id);

CREATE TABLE "rule"(
    id SERIAL,
    user_id int,
    name text
);

ALTER TABLE "rule" ADD CONSTRAINT "rule_primary_key_id" PRIMARY KEY (id);
ALTER TABLE "rule" ADD CONSTRAINT "rule_foreign_key_user_id" FOREIGN KEY (user_id) REFERENCES "user"(id);

-- +goose Down
DROP TABLE "rule";
DROP TABLE "user";
