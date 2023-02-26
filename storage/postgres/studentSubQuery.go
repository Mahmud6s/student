package postgres

import (
	"WEB-NEW-WDB/storage"
	"log"
)

const getSubjectByClassIDQuery = `SELECT * FROM subjects WHERE class=$1`

func (s PostGresDB) GetSubjectByClassID(class int) ([]storage.StSubject, error) {

	var u []storage.StSubject
	if err := s.DB.Select(&u, getSubjectByClassIDQuery, class); err != nil {
		log.Println(err)
		return nil, err
	}
	return u, nil
}

const insertMarkQuery = `
	INSERT INTO student_subjects(
		student_id,
		subject_id,
        marks
		)  
	VALUES(
		:student_id,
		:subject_id,
		:marks
		)RETURNING *;
	`

func (p PostGresDB) InsertMark(s storage.StudentSubject) (*storage.StudentSubject, error) {

	stmt, err := p.DB.PrepareNamed(insertMarkQuery)
	if err != nil {
		log.Fatalln(err)
	}

	if err := stmt.Get(&s, s); err != nil {
		log.Println(err)
		return nil, err
	}

	return &s, nil
}
