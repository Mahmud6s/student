package handler

import (
	"WEB-NEW-WDB/storage"
	"log"
	"net/http"
)

type UserForm struct {
	User        storage.User
	StClass     storage.StClass
	ListClass   []storage.StClass
	StStudent   storage.StStudent
	ListStudent []storage.StStudent
	StSubject   storage.StSubject
	ListSubject []storage.StSubject
	FormError   map[string]error
	CSRFToken   string
}

// ........////////-User-------------------------------------------
func (h Handler) CreateUserTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("create.html")
	if t == nil {
		log.Println("unable to lookup create-user template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (h Handler) ListTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("userList.html")
	if t == nil {
		log.Println("unable to lookup create-user template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (h Handler) EditUserTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("edit.html")
	if t == nil {
		log.Println("unable to lookup create-user template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (h Handler) LogInTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("login.html")
	if t == nil {
		log.Println("unable to lookup login- template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

// ..Class................////////////////////////////////////
func (h Handler) ClassCreateTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("classCreate.html")
	if t == nil {
		log.Println("unable to lookup create-class template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (h Handler) ClassListTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("classList.html")
	if t == nil {
		log.Println("unable to lookup create-class template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (h Handler) EditClassTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("classEdit.html")
	if t == nil {
		log.Println("unable to lookup create-class template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

// .................Student..////////////////////////////////////////////////////////////////
func (h Handler) StudentCreateTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("studentCreate.html")
	if t == nil {
		log.Println("unable to lookup create-student-Create template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
func (h Handler) StudentListTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("studentList.html")
	if t == nil {
		log.Println("unable to lookup create-student-List template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
func (h Handler) StudentEditTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("studentEdit.html")
	if t == nil {
		log.Println("unable to lookup create-student template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

// ........////////////////////////////////.........
func (h Handler) SubjectCreateTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("subjectCreate.html")
	if t == nil {
		log.Println("unable to lookup subject-Create template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
func (h Handler) SubjectListTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("subjectList.html")
	if t == nil {
		log.Println("unable to lookup create-subject-List template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
func (h Handler) SubjectEditTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("subjectEdit.html")
	if t == nil {
		log.Println("unable to lookup create-subject template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

// .................................................
func (h Handler) CreateMarksTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("marksSubject.html")
	if t == nil {
		log.Println("unable to lookup subject-Create template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

func (h Handler) SubjectMarksTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("subject.html")
	if t == nil {
		log.Println("unable to lookup subject-Create template")

	}
	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}

// ...
func (h Handler) ResultListTemplate(w http.ResponseWriter, data any) {
	t := h.Templates.Lookup("Result.html")
	if t == nil {
		log.Println("unable to lookup create-result-List template")

	}

	if err := t.Execute(w, data); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}
}
