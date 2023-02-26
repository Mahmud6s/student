package storage

import (
	"database/sql"
	"time"
)

type StudentSubject struct {
	ID        int `db:"id" form:"-"`
	StudentID int `db:"student_id"`
	SubjectID int `db:"subject_id"`
	Marks     int `db:"marks"`
	Mark      map[int]int
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}
