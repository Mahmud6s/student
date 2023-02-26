package postgres

import (
	"WEB-NEW-WDB/storage"
	"fmt"
	"log"
)

const insertSubjectQQ = `
INSERT INTO subjects (
	class, 
	subject1

) VALUES ( 
	:class, 
	:subject1
	
)
	returning *;
`

func (s PostGresDB) CreateSubject(u storage.StSubject) (*storage.StSubject, error) {
	stmt, err := s.DB.PrepareNamed(insertSubjectQQ)
	if err != nil {
		log.Println(err)
	}

	if err := stmt.Get(&u, u); err != nil {
		return nil, err
	}
	if u.ID == 0 {
		log.Println("Unable to Create Subject")
		return nil, fmt.Errorf("Unable to Create Subject")
	}
	return &u, nil
}

const SubjectListQQ = ` SELECT * FROM subjects WHERE deleted_at IS NULL AND (subject1 ILIKE '%%' || $1 || '%%')  order by class ASC`

func (s PostGresDB) SubjectList(uf storage.SubjectFilter) ([]storage.StSubject, error) {

	var SubjectList []storage.StSubject

	if err := s.DB.Select(&SubjectList, SubjectListQQ, uf.SearchTerm); err != nil {
		log.Println(err)
		return nil, err
	}
	return SubjectList, nil
}

const DeleteSubjectQQ = `DELETE FROM subjects WHERE id = $1 RETURNING id`

func (s PostGresDB) DeleteSubject(id string) error {
	res, err := s.DB.Exec(DeleteSubjectQQ, id)
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
		return fmt.Errorf("unable to delete subject")
	}

	return nil
}

const EditSubjectQQ = `SELECT * FROM subjects WHERE id=$1 AND deleted_at IS NULL`

func (s PostGresDB) GetSubjectIdForEdit(id string) (*storage.StSubject, error) {
	var u storage.StSubject
	if err := s.DB.Get(&u, EditSubjectQQ, id); err != nil {
		log.Println(err)
		return nil, err
	}
	return &u, nil
}

const UpdateSubjectQQ = `
	UPDATE subjects SET
	    class =:class,
		subject1=:subject1
	
	WHERE id=:id AND deleted_at IS NULL;
	`

func (s PostGresDB) UpdateSubject(u storage.StSubject) (*storage.StSubject, error) {

	stmt, err := s.DB.PrepareNamed(UpdateSubjectQQ)
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
