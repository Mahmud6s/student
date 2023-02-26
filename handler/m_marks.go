package handler

import (
	"WEB-NEW-WDB/storage"
	"net/http"

	"github.com/justinas/nosurf"
)

func (h Handler) CreateMarksR(w http.ResponseWriter, r *http.Request) {
	classList, err := h.storage.ClassList()

	if err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	studentList, err := h.storage.StudentList(storage.StudentFilter{})
	if err != nil {

		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	h.CreateMarksTemplate(w, UserForm{
		ListClass:   classList,
		ListStudent: studentList,
		CSRFToken:   nosurf.Token(r),
	})
}
