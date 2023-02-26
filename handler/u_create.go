package handler

import (
	"WEB-NEW-WDB/storage"
	"fmt"
	"log"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
)

func (h Handler) CreateUserR(w http.ResponseWriter, r *http.Request) {
	h.CreateUserTemplate(w, UserForm{
		CSRFToken: nosurf.Token(r)})
}
func (h Handler) StoreUsersR(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	//..................For-Form_Decode.............................................
	form := UserForm{}
	user := storage.User{}
	if err := h.decoder.Decode(&user, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	//..........VALIDATIN-..................
	form.User = user //again assaint-coz-if-validation is error-(Then shows error)

	if err := user.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			Nerr := make(map[string]error)
			for key, val := range vErr {
				Nerr[strings.Title(key)] = val
			}
			form.FormError = Nerr
			form.CSRFToken = nosurf.Token(r)
		}
		h.CreateUserTemplate(w, form)
		return
	}

	NewUser, err := h.storage.CreateUser(user)
	if err != nil {
		log.Fatalln(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	http.Redirect(w, r, fmt.Sprintf("/users/%v/edit", NewUser.ID), http.StatusSeeOther)
}
