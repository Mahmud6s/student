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

func (h Handler) EditSubjectR(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	//......
	editSubject, err := h.storage.GetSubjectIdForEdit(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	var form UserForm
	form.StSubject = *editSubject
	form.CSRFToken = nosurf.Token(r)
	h.SubjectEditTemplate(w, form)
}

func (h Handler) UpdateSubjectR(w http.ResponseWriter, r *http.Request) {

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
	subject := storage.StSubject{ID: uID}
	if err := h.decoder.Decode(&subject, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	//.............................................................
	form.StSubject = subject
	if err := subject.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			Nerr := make(map[string]error)
			for key, val := range vErr {
				Nerr[strings.Title(key)] = val
			}
			form.FormError = Nerr
			form.CSRFToken = nosurf.Token(r)
		}
		h.SubjectEditTemplate(w, form)
		return
	}

	// updateStorage, err := h.storage.UpdateUser()
	h.storage.UpdateSubject(subject)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

	http.Redirect(w, r, "/subject/list", http.StatusSeeOther)

}
