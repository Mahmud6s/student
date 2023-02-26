package handler

import (
	"WEB-NEW-WDB/storage"
	"log"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
)

func (h Handler) CreateStudentR(w http.ResponseWriter, r *http.Request) {
	classList, err := h.storage.ClassList()
	if err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	h.StudentCreateTemplate(w, UserForm{
		ListClass: classList,
		CSRFToken: nosurf.Token(r),
	})
}
func (h Handler) StoreStudentR(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	//..................For-Form_Decode.............................................

	form := UserForm{}

	student := storage.StStudent{}
	if err := h.decoder.Decode(&student, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	//..........VALIDATIN-..................

	form.StStudent = student //again assaint-coz-if-validation is error-(Then shows error)

	if err := student.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			Nerr := make(map[string]error)
			for key, val := range vErr {
				Nerr[strings.Title(key)] = val
			}
			form.FormError = Nerr
			form.CSRFToken = nosurf.Token(r)
		}
		h.StudentCreateTemplate(w, form)
		return
	}

	// NewClass, err :=
	data, err := h.storage.CreateStudent(student)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	h.MarksHandler(w, r, student.StudentClass, data.ID)
	http.Redirect(w, r, "/student/list", http.StatusSeeOther)
}
