package postgres

import (
	"WEB-NEW-WDB/storage"
	"fmt"
	"log"

	"golang.org/x/crypto/bcrypt"
)

// ..............SHOW-UserList.........................................
const ListQQ = `SELECT * FROM users WHERE deleted_at IS NULL AND (first_name ILIKE '%%' || $1 || '%%' OR last_name ILIKE '%%' || $1 || '%%' OR username ILIKE '%%' || $1 || '%%' OR email ILIKE '%%' || $1 || '%%')  order by id ASC`

// .`SELECT * FROM users WHERE deleted_at IS NULL order by id asc`
func (s PostGresDB) ListUser(uf storage.UserFilter) ([]storage.User, error) {
	var UserList []storage.User
	if err := s.DB.Select(&UserList, ListQQ, uf.SearchTerm); err != nil {
		log.Println(err)
		return nil, err
	}
	return UserList, nil
}

// ..............SHOW-UserList(end).........................................

// ................CREATE-USER(start)..........................................
const insertQQ = `
	INSERT INTO users(
		first_name,
		last_name,
		username,
		email,
		password
	) VALUES (
		:first_name,
		:last_name,
		:username,
		:email,
		:password
	) RETURNING *;
`

func (s PostGresDB) CreateUser(u storage.User) (*storage.User, error) {
	// var user storage.User
	stmt, err := s.DB.PrepareNamed(insertQQ)
	if err != nil {
		log.Println(err)
	}
	//................HassPss(start)..................................................
	HassPass, err := bcrypt.GenerateFromPassword([]byte(u.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, err
	}
	u.Password = string(HassPass)
	//................HassPss(end)..................................................
	if err := stmt.Get(&u, u); err != nil {
		return nil, err
	}
	if u.ID == 0 {
		log.Println("Unable to Create User")
		return nil, fmt.Errorf("Unable to Create User")
	}
	return &u, nil
}

// ................CREATE-USER(end)..........................................

// ............User-Edit/Update(start).........................................
const EditQQ = `SELECT * FROM users WHERE id=$1 AND deleted_at IS NULL`

func (s PostGresDB) GetUserIdForEdit(id string) (*storage.User, error) {
	var u storage.User
	if err := s.DB.Get(&u, EditQQ, id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

const UpdateQQ = `
	UPDATE users SET
	    first_name =:first_name,
		last_name =:last_name,
		status =:status
	WHERE id=:id AND deleted_at IS NULL;
	`

func (s PostGresDB) UpdateUser(u storage.User) (*storage.User, error) {

	stmt, err := s.DB.PrepareNamed(UpdateQQ)
	if err != nil {
		log.Fatalln(err)
	}

	res, err := stmt.Exec(u)
	if err != nil {
		log.Fatalln(err)
	}
	Rcount, err := res.RowsAffected()
	if Rcount < 1 || err != nil {
		log.Fatalln(err)
	}
	return &u, nil

}

// ............User-Edit/Update(end)...................................

// ..............DELETE-Start(start)---..................................
// const DeleteIdQQ = `SELECT * FROM users WHERE id=$1 AND deleted_at IS NULL`

// func (s PostGresDB) GetUserIdForDelete(id string) (*storage.User, error) {
// 	var u storage.User
// 	if err := s.DB.Get(&u, DeleteIdQQ, id); err != nil {
// 		log.Println(err)
// 		return nil, err
// 	}

// 	return &u, nil
// }

// `UPDATE users SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`
// `DELETE FROM users WHERE id = $1 RETURNING id`
const deleteQQ = `DELETE FROM users WHERE id = $1 RETURNING id`

func (s PostGresDB) DeleteUser(id string) error {
	res, err := s.DB.Exec(deleteQQ, id)
	if err != nil {
		fmt.Println(err)
		return err
	}

	rowCount, err := res.RowsAffected()
	if err != nil {
		fmt.Println(err)
		return err
	}

	if rowCount <= 0 {
		return fmt.Errorf("unable to delete user")
	}

	return nil
}

// ..............DELETE-(end)---..................................

const GetUsernameQQ = `SELECT * FROM users WHERE username=$1 AND deleted_at IS NULL`

func (s PostGresDB) GetUsernameForLogin(username string) (*storage.User, error) {
	var u storage.User
	if err := s.DB.Get(&u, GetUsernameQQ, username); err != nil {
		log.Println(err)
		return nil, err
	}

	return &u, nil
}

// ..
