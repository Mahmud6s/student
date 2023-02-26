package handler

import (
	"WEB-NEW-WDB/storage"
	"log"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
)

func (h Handler) CreateSubjectR(w http.ResponseWriter, r *http.Request) {
	classList, err := h.storage.ClassList()
	if err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	h.SubjectCreateTemplate(w, UserForm{
		ListClass: classList,
		CSRFToken: nosurf.Token(r),
	})
}
func (h Handler) StoreSubjectR(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	//..................For-Form_Decode.............................................

	form := UserForm{}

	subject := storage.StSubject{}
	if err := h.decoder.Decode(&subject, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	//..........VALIDATIN-..................

	form.StSubject = subject //again assaint-coz-if-validation is error-(Then shows error)

	if err := subject.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			Nerr := make(map[string]error)
			for key, val := range vErr {
				Nerr[strings.Title(key)] = val
			}
			form.FormError = Nerr
			form.CSRFToken = nosurf.Token(r)
		}
		h.SubjectCreateTemplate(w, form)
		return
	}

	// NewClass, err :=
	_, err := h.storage.CreateSubject(subject)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	http.Redirect(w, r, "/subject/list", http.StatusSeeOther)
}
