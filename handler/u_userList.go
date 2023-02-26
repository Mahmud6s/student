package handler

import (
	"WEB-NEW-WDB/storage"
	"net/http"
)

type SearchForm struct {
	User       []storage.User
	StStudent  []storage.StStudent
	StSubject  []storage.StSubject
	SearchTerm string
}

func (h Handler) ListUserR(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	st := r.FormValue("SearchTerm")
	uf := storage.UserFilter{
		SearchTerm: st,
	}

	listUser, err := h.storage.ListUser(uf)
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	data := SearchForm{
		User:       listUser,
		SearchTerm: st,
	}

	h.ListTemplate(w, data)
}
