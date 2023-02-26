package storage

import (
	"database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type StudentFilter struct {
	SearchTerm string
}
type StStudent struct {
	ID           int          `db:"id" form:"-" `
	FirstName    string       `db:"first_name"`
	LastName     string       `db:"last_name"`
	StudentClass int          `db:"student_class"`
	StudentRoll  int          `db:"student_roll"`
	CreatedAt    time.Time    `db:"created_at"`
	UpdatedAt    time.Time    `db:"updated_at"`
	DeletedAt    sql.NullTime `db:"deleted_at"`
}

func (ss StStudent) Validate() error {
	return validation.ValidateStruct(&ss,
		validation.Field(&ss.FirstName,
			validation.Required.Error("The first name field is required."),
		),
		validation.Field(&ss.LastName,
			validation.Required.Error("The last name field is required."),
		),
		validation.Field(&ss.StudentClass,
			validation.Required.Error("The Student Class field is required."),
		),
		validation.Field(&ss.StudentRoll,
			validation.Required.Error("The Student Roall field is required."),
			validation.Required, validation.Min(1), validation.Max(100).Error(" must input 1 to 100 "),
		),
	)
}
