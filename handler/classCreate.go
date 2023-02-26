package handler

import (
	"WEB-NEW-WDB/storage"
	"log"
	"net/http"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
)

func (h Handler) CreateClassR(w http.ResponseWriter, r *http.Request) {
	h.ClassCreateTemplate(w, UserForm{
		CSRFToken: nosurf.Token(r),
	})
}
func (h Handler) StoreClassR(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	//..................For-Form_Decode.............................................
	form := UserForm{}
	class := storage.StClass{}
	if err := h.decoder.Decode(&class, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	//..........VALIDATIN-..................
	form.StClass = class //again assaint-coz-if-validation is error-(Then shows error)

	if err := class.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			Nerr := make(map[string]error)
			for key, val := range vErr {
				Nerr[strings.Title(key)] = val
			}
			form.FormError = Nerr
			form.CSRFToken = nosurf.Token(r)
		}
		h.ClassCreateTemplate(w, form)
		return
	}

	// NewClass, err :=
	h.storage.CreateClass(class)
	// if err != nil {
	// 	log.Fatalln(err)
	// 	http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	// }

	http.Redirect(w, r, "/class/list", http.StatusSeeOther)
}
