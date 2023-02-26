package handler

import (
	"WEB-NEW-WDB/storage"
	"net/http"
)

func (h Handler) StudentListR(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	st := r.FormValue("SearchTerm")
	uf := storage.StudentFilter{
		SearchTerm: st,
	}
	studentlist, err := h.storage.StudentList(storage.StudentFilter(uf))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	data := SearchForm{
		StStudent:  studentlist,
		SearchTerm: st,
	}
	h.StudentListTemplate(w, data)

}
