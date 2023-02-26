package handler

import (
	"WEB-NEW-WDB/storage"
	"log"
	"net/http"
	"strconv"
	"strings"

	"github.com/go-chi/chi"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	_ "github.com/lib/pq"
)

func (h Handler) EditStudentR(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	//......
	editStudent, err := h.storage.GetStudentIdForEdit(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	var form UserForm
	form.StStudent = *editStudent
	form.CSRFToken = nosurf.Token(r)
	h.StudentEditTemplate(w, form)
}

func (h Handler) UpdateStudentR(w http.ResponseWriter, r *http.Request) {

	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	//.....
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	// //.....................................................
	var form UserForm
	student := storage.StStudent{ID: uID}
	if err := h.decoder.Decode(&student, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	//.............................................................
	form.StStudent = student
	if err := student.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			Nerr := make(map[string]error)
			for key, val := range vErr {
				Nerr[strings.Title(key)] = val
			}
			form.FormError = Nerr
			form.CSRFToken = nosurf.Token(r)
		}
		h.StudentEditTemplate(w, form)
		return
	}

	// updateStorage, err := h.storage.UpdateUser()
	h.storage.UpdateStudent(student)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

	http.Redirect(w, r, "/student/list", http.StatusSeeOther)

}
