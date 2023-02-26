package storage

import (
	"database/sql"
	"time"
)

type Mark struct {
	Class   string
	Student string
}
type InputStore struct {
	ID          int            `db:"id" form:"id"`
	FirstName   string         `db:"first_name" form:"first_name"`
	LastName    string         `db:"last_name" form:"last_name"`
	Class       sql.NullInt64  `db:"class" form:"class"`
	StudentRoll int            `db:"student_roll" form:"student_roll"`
	SubjectID   int            `db:"subject_id" form:"subject_id"`
	Subject1    sql.NullString `db:"subject1" form:"subject1"`
	CreatedAt   time.Time      `db:"created_at" form:"created_at"`
	UpdatedAt   time.Time      `db:"updated_at" form:"updated_at"`
	DeletedAt   sql.NullTime   `db:"deleted_at" form:"deleted_at"`
}
type Result struct {
	ID          int            `db:"id" form:"id"`
	FirstName   string         `db:"first_name" form:"first_name"`
	LastName    string         `db:"last_name" form:"last_name"`
	Class       sql.NullInt64  `db:"class" form:"class"`
	StudentRoll int            `db:"student_roll" form:"student_roll"`
	Subject1    sql.NullString `db:"subject1" form:"subject1"`
	Marks       int            `db:"marks"`
}
