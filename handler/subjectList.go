package handler

import (
	"WEB-NEW-WDB/storage"
	"net/http"
)

func (h Handler) ListSubjectR(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	st := r.FormValue("SearchTerm")
	uf := storage.UserFilter{
		SearchTerm: st,
	}
	subjectlist, err := h.storage.SubjectList(storage.SubjectFilter(uf))
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	data := SearchForm{
		StSubject:  subjectlist,
		SearchTerm: st,
	}

	h.SubjectListTemplate(w, data)

}
