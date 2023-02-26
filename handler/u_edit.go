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

func (h Handler) EditUserR(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	//......
	editUser, err := h.storage.GetUserIdForEdit(id)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	var form UserForm
	form.User = *editUser
	form.CSRFToken = nosurf.Token(r)
	h.EditUserTemplate(w, form)
}

func (h Handler) UpdateUserR(w http.ResponseWriter, r *http.Request) {
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
	user := storage.User{ID: uID}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	//.............................................................
	form.User = user
	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			Nerr := make(map[string]error)
			for key, val := range vErr {
				Nerr[strings.Title(key)] = val
			}
			form.FormError = Nerr
			form.CSRFToken = nosurf.Token(r)
		}
		h.EditUserTemplate(w, form)
		return
	}

	// updateStorage, err := h.storage.UpdateUser(user)
	h.storage.UpdateUser(user)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

	http.Redirect(w, r, "/users", http.StatusSeeOther)

}
