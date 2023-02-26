package postgres

import (
	"WEB-NEW-WDB/storage"
	"fmt"
	"log"
)

const insertStudentQQ = `
	INSERT INTO students(
		first_name,
		last_name,
		student_class,
		student_roll
	) VALUES (
		:first_name,
		:last_name,
		:student_class,
		:student_roll	
	) RETURNING *;
`

func (s PostGresDB) CreateStudent(u storage.StStudent) (*storage.StStudent, error) {

	stmt, err := s.DB.PrepareNamed(insertStudentQQ)
	if err != nil {
		log.Println(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		return nil, err
	}
	if u.ID == 0 {
		log.Println("Unable to Create student")
		return nil, fmt.Errorf("Unable to Create student")
	}
	return &u, nil
}

const StudentListQQ = ` SELECT * FROM students WHERE deleted_at IS NULL AND (first_name ILIKE '%%' || $1 || '%%' OR last_name ILIKE '%%' || $1 || '%%')  order by student_class ASC`

// `SELECT students.first_name,students.last_name,students.roll, stuclass.class_name
// FROM students
// INNER JOIN stuclass ON students.student_class = stuclass.id;`

func (s PostGresDB) StudentList(uf storage.StudentFilter) ([]storage.StStudent, error) {

	var StudentList []storage.StStudent

	if err := s.DB.Select(&StudentList, StudentListQQ, uf.SearchTerm); err != nil {
		log.Println(err)
		return nil, err
	}
	return StudentList, nil
}

// ..Delete
const DeleteStudentQQ = `DELETE FROM students WHERE id = $1 RETURNING id`

func (s PostGresDB) DeleteStudent(id string) error {
	res, err := s.DB.Exec(DeleteStudentQQ, id)
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
		return fmt.Errorf("unable to delete student")
	}

	return nil
}

// ..
const EditStudentQQ = `SELECT * FROM students WHERE id=$1 AND deleted_at IS NULL`

func (s PostGresDB) GetStudentIdForEdit(id string) (*storage.StStudent, error) {
	var u storage.StStudent
	if err := s.DB.Get(&u, EditStudentQQ, id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

const UpdateStudentQQ = `
	UPDATE students SET
	    first_name =:first_name,
		last_name =:last_name,
		student_class =:student_class,
		student_roll=:student_roll
		
	WHERE id=:id AND deleted_at IS NULL;
	`

func (s PostGresDB) UpdateStudent(u storage.StStudent) (*storage.StStudent, error) {

	stmt, err := s.DB.PrepareNamed(UpdateStudentQQ)
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
