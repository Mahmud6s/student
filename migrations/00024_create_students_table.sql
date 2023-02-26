-- +goose Up
-- +goose StatementBegin
CREATE TABLE IF NOT EXISTS students (
    id BIGSERIAL PRIMARY KEY ,
    first_name TEXT NOT NULL,
    last_name TEXT NOT NULL,
    student_class INT NOT NULL,
    student_roll TEXT NOT NULL,
	created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
	deleted_at TIMESTAMP DEFAULT NULL

);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS students;
-- +goose StatementEnd
