package handler

import (
	"WEB-NEW-WDB/storage/postgres"
	"fmt"
	"log"
	"net/http"
	"regexp"
	"strconv"
	"strings"

	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/justinas/nosurf"
	"golang.org/x/crypto/bcrypt"
)

type LoginUser struct {
	Username  string
	Password  string
	FormError map[string]error
	CSRFToken string
}

func (lu LoginUser) Validate() error {
	return validation.ValidateStruct(&lu,
		validation.Field(&lu.Username,
			validation.Required.Error("The username field is required."),
			validation.Match(regexp.MustCompile(`^\S+$`)).Error("username cannot contain spaces"),
			validation.Required, validation.Length(6, 20),
		),
		validation.Field(&lu.Password,
			validation.Required.Error("The password field is required."),
			validation.Length(6, 12).Error("filed is 6 to 12 numbers"),
			validation.Match(regexp.MustCompile(`^\S+$`)).Error("password cannot contain spaces"),
		),
	)

}

func (h Handler) Login(w http.ResponseWriter, r *http.Request) {
	h.LogInTemplate(w, LoginUser{
		CSRFToken: nosurf.Token(r),
	})
}

func (h Handler) LoginUserR(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	var lf LoginUser
	if err := h.decoder.Decode(&lf, r.PostForm); err != nil {
		log.Println(err)
		http.Error(w, "internal server error", http.StatusInternalServerError)
	}

	if err := lf.Validate(); err != nil {
		if vErr, ok := err.(validation.Errors); ok {
			formErr := make(map[string]error)
			for key, val := range vErr {
				formErr[strings.Title(key)] = val
			}
			lf.FormError = formErr
			lf.Password = ""
			lf.CSRFToken = nosurf.Token(r)
			h.LogInTemplate(w, lf)
			return
		}
	}

	user, err := h.storage.GetUsernameForLogin(lf.Username)
	if err != nil {
		if err.Error() == postgres.NotFound {
			formErr := make(map[string]error)
			formErr["Username"] = fmt.Errorf("credentials does not match")
			lf.FormError = formErr
			lf.CSRFToken = nosurf.Token(r)
			lf.Password = ""
			h.LogInTemplate(w, lf)
			return
		}

		http.Error(w, "internal server error", http.StatusInternalServerError)
		return
	}

	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(lf.Password)); err != nil {
		formErr := make(map[string]error)
		formErr["Username"] = fmt.Errorf("credentials does not match")
		lf.FormError = formErr
		lf.CSRFToken = nosurf.Token(r)
		lf.Password = ""
		h.LogInTemplate(w, lf)
		return
	}

	h.sessionManager.Put(r.Context(), "userID", strconv.Itoa(user.ID))
	http.Redirect(w, r, "/users", http.StatusSeeOther)
}
