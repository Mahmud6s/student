package postgres

import (
	"WEB-NEW-WDB/storage"
	"fmt"
	"log"
)

const insertClassQQ = `
	INSERT INTO stuclass(
		class_name
	
	) VALUES (
		:class_name
		
	) RETURNING *;
`

func (s PostGresDB) CreateClass(st storage.StClass) (*storage.StClass, error) {
	// var user storage.User
	stmt, err := s.DB.PrepareNamed(insertClassQQ)
	if err != nil {
		log.Println(err)
	}

	if err := stmt.Get(&st, st); err != nil {
		return nil, err
	}
	if st.ID == 0 {
		log.Println("Unable to Create Class")
		return nil, fmt.Errorf("Unable to Create Class")
	}
	return &st, nil
}

// ..
const ClassListQQ = `SELECT * FROM stuclass WHERE deleted_at IS NULL order by id asc`

func (s PostGresDB) ClassList() ([]storage.StClass, error) {

	var ClassList []storage.StClass

	if err := s.DB.Select(&ClassList, ClassListQQ); err != nil {
		log.Println(err)
		return nil, err
	}
	return ClassList, nil
}

// ..classDELETE

// `UPDATE stuclass SET deleted_at = CURRENT_TIMESTAMP WHERE id=$1 AND deleted_at IS NULL`
// `DELETE FROM stuclass WHERE id = $1 RETURNING id`
const deleteClassQQ = `DELETE FROM stuclass WHERE id = $1 RETURNING id`

func (s PostGresDB) DeleteClass(id string) error {
	res, err := s.DB.Exec(deleteClassQQ, id)
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

// ..Edite-&-Update
const EditClassQQ = `SELECT * FROM stuclass WHERE id=$1 AND deleted_at IS NULL`

func (s PostGresDB) GetClassIdForEdit(id string) (*storage.StClass, error) {
	var u storage.StClass
	if err := s.DB.Get(&u, EditClassQQ, id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

const UpdateClassQQ = `
	UPDATE stuclass SET
	    class_name =:class_name
		
	WHERE id=:id AND deleted_at IS NULL;
	`

func (s PostGresDB) UpdateClass(u storage.StClass) (*storage.StClass, error) {

	stmt, err := s.DB.PrepareNamed(UpdateClassQQ)
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
