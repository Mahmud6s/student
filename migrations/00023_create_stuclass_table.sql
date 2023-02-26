-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS stuclass (
    id BIGSERIAL,
    class_name TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP DEFAULT NULL,

	PRIMARY KEY(id)
	
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS stuclass;
-- +goose StatementEnd
