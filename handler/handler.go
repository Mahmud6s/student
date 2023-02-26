package handler

import (
	"WEB-NEW-WDB/storage"
	"log"
	"net/http"
	"os"
	"text/template"

	"github.com/alexedwards/scs/v2"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-playground/form"
)

type Handler struct {
	decoder        *form.Decoder
	sessionManager *scs.SessionManager
	storage        dbstorege
	Templates      *template.Template
}

type dbstorege interface {
	ListUser(storage.UserFilter) ([]storage.User, error)
	CreateUser(u storage.User) (*storage.User, error)
	GetUserIdForEdit(id string) (*storage.User, error)
	UpdateUser(u storage.User) (*storage.User, error)
	DeleteUser(id string) error
	// GetUserIdForDelete(id string) (*storage.User, error)
	GetUsernameForLogin(username string) (*storage.User, error)
	//..class()
	CreateClass(st storage.StClass) (*storage.StClass, error)
	ClassList() ([]storage.StClass, error)
	DeleteClass(id string) error
	UpdateClass(u storage.StClass) (*storage.StClass, error)
	GetClassIdForEdit(id string) (*storage.StClass, error)
	//....Students()
	CreateStudent(u storage.StStudent) (*storage.StStudent, error)
	StudentList(storage.StudentFilter) ([]storage.StStudent, error)
	DeleteStudent(id string) error
	GetStudentIdForEdit(id string) (*storage.StStudent, error)
	UpdateStudent(u storage.StStudent) (*storage.StStudent, error)
	//...subjects
	CreateSubject(u storage.StSubject) (*storage.StSubject, error)
	SubjectList(storage.SubjectFilter) ([]storage.StSubject, error)
	DeleteSubject(id string) error
	GetSubjectIdForEdit(id string) (*storage.StSubject, error)
	UpdateSubject(u storage.StSubject) (*storage.StSubject, error)

	GetSubjectByClassID(class int) ([]storage.StSubject, error)

	InsertMark(s storage.StudentSubject) (*storage.StudentSubject, error)
	//.....
	GetSubject(id string) ([]storage.InputStore, error)
	GetMark(u storage.StudentSubject) (*storage.StudentSubject, error)
	ResultListt(id int) ([]storage.Result, error)
	//..
	// DeleteResult(id string) error
}

// ................................................................
func NewHandler(formDecoder *form.Decoder, sm *scs.SessionManager, storage dbstorege) *chi.Mux {
	h := &Handler{
		decoder:        formDecoder,
		sessionManager: sm,
		storage:        storage,
	}

	h.All_TemplateParse() //
	r := chi.NewRouter()
	//..............MIDDLE-WARE..................................................
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	//.....................................................................
	//..............ROUTE(start)...........................................
	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Get("/", h.Home)
		r.Get("/login", h.Login)
		r.Post("/login", h.LoginUserR)
		r.Get("/users/create", h.CreateUserR)
		r.Post("/users/store", h.StoreUsersR)
	})

	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Get("/logout", h.Logout)

		r.Route("/users", func(r chi.Router) {
			r.Get("/", h.ListUserR)
			r.Get("/{id:[0-9]+}/edit", h.EditUserR)
			r.Post("/{id:[0-9]+}/update", h.UpdateUserR)
			r.Get("/{id:[0-9]+}/delete", h.DeleteUserR)
		})
	})

	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Route("/class", func(r chi.Router) {
			r.Get("/create", h.CreateClassR)
			r.Post("/store", h.StoreClassR)
			r.Get("/list", h.ClassListR)
			r.Get("/{id:[0-9]+}/edit", h.EditClassR)
			r.Post("/{id:[0-9]+}/update", h.UpdateClassR)
			r.Get("/{id:[0-9]+}/delete", h.DeleteClassR)
		})
	})
	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Route("/student", func(r chi.Router) {
			r.Get("/create", h.CreateStudentR)
			r.Post("/store", h.StoreStudentR)
			r.Get("/list", h.StudentListR)
			r.Get("/{id:[0-9]+}/edit", h.EditStudentR)
			r.Post("/{id:[0-9]+}/update", h.UpdateStudentR)
			r.Get("/{id:[0-9]+}/delete", h.DeleteStudentR)
		})
	})
	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Route("/subject", func(r chi.Router) {
			r.Get("/create", h.CreateSubjectR)
			r.Post("/store", h.StoreSubjectR)
			r.Get("/list", h.ListSubjectR)
			r.Get("/{id:[0-9]+}/edit", h.EditSubjectR)
			r.Post("/{id:[0-9]+}/update", h.UpdateSubjectR)
			r.Get("/{id:[0-9]+}/delete", h.DeleteSubjectR)
		})
	})
	r.Group(func(r chi.Router) {
		r.Use(sm.LoadAndSave)
		r.Use(h.Authentication)
		r.Route("/marks", func(r chi.Router) {
			r.Get("/create", h.CreateMarksR)
			r.Post("/subject", h.SubjectMarkR)
			r.Post("/store", h.MarksInputR)
			r.Get("/{id:[0-9]+}/result", h.ResultListR)
			// r.Post("/{id:[0-9]+}/update", h.UpdateSubjectR)
			// r.Get("/{id:[0-9]+}/delete", h.DeleteResultR)
		})
	})
	return r
}

// ..............ROUTE(end)...........................................
// ..............Auth of login............................................
func (h Handler) Authentication(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		user_ID := h.sessionManager.GetString(r.Context(), "userID")
		if user_ID == "" {
			http.Redirect(w, r, "/login", http.StatusSeeOther)
			return
		}
		next.ServeHTTP(w, r)
	})
}

// ......................................................................
func (h *Handler) All_TemplateParse() error {
	templates := template.New("std-templates").Funcs(template.FuncMap{
		"globalFunc": func(n string) string {
			return ""
		},
	})
	NewFs := os.DirFS("assets/templates")
	tmpl := template.Must(templates.ParseFS(NewFs, "*/*/*.html", "*.html"))
	if tmpl == nil {
		log.Fatalln("Unable to parse Template file")
	}
	h.Templates = tmpl
	return nil
}

//.................................................................
