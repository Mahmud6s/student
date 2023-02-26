package handler

// import (
// 	"net/http"

// 	"github.com/go-chi/chi"
// )

// func (h Handler) DeleteResultR(w http.ResponseWriter, r *http.Request) {
// 	id := chi.URLParam(r, "id")
// 	if err := h.storage.DeleteResult(id); err != nil {
// 		http.Error(w, "internal server error", http.StatusInternalServerError)
// 		return
// 	}
// 	http.Redirect(w, r, "/student/list", http.StatusPermanentRedirect)
// }
