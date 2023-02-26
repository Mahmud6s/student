package handler

import (
	"WEB-NEW-WDB/storage"
	"log"
	"net/http"

	"github.com/justinas/nosurf"
)

type MarkForm struct {
	SubList   []storage.InputStore
	Class     string
	Student   string
	CSRFToken string
}

func (h Handler) SubjectMarkR(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	var form MarkForm

	if err := h.decoder.Decode(&form, r.PostForm); err != nil {
		log.Fatalln(err)

	}

	SubList, err := h.storage.GetSubject(form.Student)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}
	h.SubjectMarksTemplate(w, MarkForm{
		SubList:   SubList,
		CSRFToken: nosurf.Token(r),
	})
}

func (h Handler) MarksInputR(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)

	}

	marks := storage.StudentSubject{}
	if err := h.decoder.Decode(&marks, r.PostForm); err != nil {
		log.Fatalln(err)
	}

	for id, mark := range marks.Mark {

		m := storage.StudentSubject{
			ID:    id,
			Marks: mark,
		}
		_, err := h.storage.GetMark(m)
		if err != nil {
			log.Println(err)
			http.Error(w, "internal server error", http.StatusInternalServerError)
		}
	}

	http.Redirect(w, r, "/marks/create", http.StatusSeeOther)
}
