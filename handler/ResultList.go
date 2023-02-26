package handler

import (
	"WEB-NEW-WDB/storage"
	"log"
	"net/http"
	"strconv"

	"github.com/go-chi/chi"
)

type ResultForm struct {
	SelectedStu  []storage.Result
	SelectedMark storage.Result
}

func (h Handler) ResultListR(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	uID, err := strconv.Atoi(id)
	if err != nil {
		log.Println(err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}

	resultList, err := h.storage.ResultListt(uID)

	if err != nil {
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
	data := ResultForm{
		SelectedStu:  resultList,
		SelectedMark: resultList[0],
	}
	h.ResultListTemplate(w, data)
}
