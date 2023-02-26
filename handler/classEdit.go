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

func (h Handler) EditClassR(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	//......
	editClass, err := h.storage.GetClassIdForEdit(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	var form UserForm
	form.StClass = *editClass
	form.CSRFToken = nosurf.Token(r)
	h.EditClassTemplate(w, form)
}

func (h Handler) UpdateClassR(w http.ResponseWriter, r *http.Request) {
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
	class := storage.StClass{ID: uID}
	if err := h.decoder.Decode(&class, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	//.............................................................
	form.StClass = class
	if err := class.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			Nerr := make(map[string]error)
			for key, val := range vErr {
				Nerr[strings.Title(key)] = val
			}
			form.FormError = Nerr
			form.CSRFToken = nosurf.Token(r)
		}
		h.EditClassTemplate(w, form)
		return
	}

	// updateStorage, err := h.storage.UpdateUser()
	h.storage.UpdateClass(class)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

	http.Redirect(w, r, "/class/list", http.StatusSeeOther)

}
