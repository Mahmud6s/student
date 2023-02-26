package storage

import (
	"database/sql"
	"time"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type SubjectFilter struct {
	SearchTerm string
}
type StSubject struct {
	ID        int          `db:"id" form:"-" `
	Class     int          `db:"class"`
	Subject1  string       `db:"subject1"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (ss StSubject) Validate() error {
	return validation.ValidateStruct(&ss,
		validation.Field(&ss.Class,
			validation.Required.Error("The Class field is required."),
		),
		validation.Field(&ss.Subject1,
			validation.Required.Error("The Subject1 field is required."),
		),
	)
}

// func UserEmailUnique(value interface{}) error {
// 	email := value.(string)
// 	var count int
// 	err := DB.QueryRow("SELECT count(*) FROM users WHERE email = $1", email).Scan(&count)
// 	if err != nil {
// 		// Handle error
// 	}
// 	if count > 0 {
// 		return validation.NewError("Email already exists")
// 	}
// 	return nil
// }
