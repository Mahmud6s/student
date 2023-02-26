package storage

import (
	"database/sql"
	"regexp"
	"time"

	"github.com/go-ozzo/ozzo-validation/is"
	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UserFilter struct {
	SearchTerm string
}

type User struct {
	ID        int          `db:"id" form:"-" `
	FirstName string       `db:"first_name"`
	LastName  string       `db:"last_name"`
	Email     string       `db:"email"`
	Username  string       `db:"username"`
	Password  string       `db:"password"`
	Status    bool         `db:"status"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (u User) Validate() error {
	return validation.ValidateStruct(&u,
		validation.Field(&u.FirstName,
			validation.Required.Error("The first name field is required."),
		),
		validation.Field(&u.LastName,
			validation.Required.Error("The last name field is required."),
		),
		validation.Field(&u.Username,
			validation.Required.Error("The username field is required."),
			validation.Match(regexp.MustCompile(`^\S+$`)).Error("username cannot contain spaces"),
			validation.Required, validation.Length(6, 20),
		),
		validation.Field(&u.Email,
			validation.Required.When(u.ID == 0).Error("The email field is required."),
			is.Email.Error("The email field must be a valid email."),
		),
		validation.Field(&u.Password,
			validation.Required.When(u.ID == 0).Error("The password field is required."),
			validation.Length(6, 12).Error("filed is 6 to 12 number"),
			validation.Match(regexp.MustCompile(`^\S+$`)).Error("password cannot contain spaces"),
		),
	)
}

type StClass struct {
	ID        int          `db:"id" form:"-"`
	ClassName string       `db:"class_name"`
	CreatedAt time.Time    `db:"created_at"`
	UpdatedAt time.Time    `db:"updated_at"`
	DeletedAt sql.NullTime `db:"deleted_at"`
}

func (st StClass) Validate() error {
	return validation.ValidateStruct(&st,
		validation.Field(&st.ClassName,
			validation.Required.Error("The Class field is required."),
		),
		validation.Field(&st.ClassName,
			validation.Match(regexp.MustCompile(`^Class [1-9]$|^Class 10$`)).Error("Class must be in the format 'Class [1-10]'"),
		),
	)
}
