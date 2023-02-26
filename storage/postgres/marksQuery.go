package postgres

import (
	"WEB-NEW-WDB/storage"
	"fmt"
	"log"
)

const GetSubjectQQ = `SELECT subjects.subject1, subjects.class, students.first_name, students.last_name, students.student_roll,student_subjects.subject_id,student_subjects.id
FROM subjects
FULL OUTER JOIN student_subjects ON subjects.id = student_subjects.subject_id
FULL OUTER JOIN students ON students.id = student_subjects.student_id
WHERE students.id = $1
ORDER BY subjects.subject1;`

func (s PostGresDB) GetSubject(id string) ([]storage.InputStore, error) {
	var u []storage.InputStore
	if err := s.DB.Select(&u, GetSubjectQQ, id); err != nil {
		log.Fatalln(err)
		return nil, err
	}
	return u, nil
}

const IncertMarksQQ = `
UPDATE student_subjects
SET marks = :marks
WHERE id = :id
	returning *;`

func (s PostGresDB) GetMark(u storage.StudentSubject) (*storage.StudentSubject, error) {

	stmt, err := s.DB.PrepareNamed(IncertMarksQQ)
	if err != nil {
		log.Println(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		return nil, err
	}
	if u.ID == 0 {
		log.Println("Unable to Create marks")
		return nil, fmt.Errorf("Unable to Create Marks")
	}
	return &u, nil
}

const ResultListQQ = `SELECT subjects.subject1, subjects.class, students.first_name, students.last_name, students.student_roll,student_subjects.id,student_subjects.marks
FROM subjects
FULL OUTER JOIN student_subjects ON subjects.id = student_subjects.subject_id
FULL OUTER JOIN students ON students.id = student_subjects.student_id
WHERE student_subjects.student_id = $1
`

func (s PostGresDB) ResultListt(id int) ([]storage.Result, error) {

	var ResulList []storage.Result

	if err := s.DB.Select(&ResulList, ResultListQQ, id); err != nil {
		log.Println(err)
		return nil, err
	}
	return ResulList, nil
}

// ...
// const DeleteResultQQ = `DELETE FROM student_subjects WHERE id = $1 RETURNING id`

// func (s PostGresDB) DeleteResult(id string) error {
// 	res, err := s.DB.Exec(DeleteResultQQ, id)
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	rowCount, err := res.RowsAffected()
// 	if err != nil {
// 		fmt.Println(err)
// 		return err
// 	}

// 	if rowCount <= 0 {
// 		return fmt.Errorf("unable to delete result")
// 	}

// 	return nil
// }
