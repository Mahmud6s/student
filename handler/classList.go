package handler

import (
	"net/http"
)

func (h Handler) ClassListR(w http.ResponseWriter, r *http.Request) {

	classlist, err := h.storage.ClassList()
	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	h.ClassListTemplate(w, classlist)

}
