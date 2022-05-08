package router

import (
	"github.com/go-chi/chi/v5"
	"net/http"
)

type Handler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	CreateRecord(w http.ResponseWriter, r *http.Request)
	ReadRecord(w http.ResponseWriter, r *http.Request)
	UpdateRecord(w http.ResponseWriter, r *http.Request)
	DeleteRecord(w http.ResponseWriter, r *http.Request)
}

func NewRouter(uh, ah Handler) http.Handler {
	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("REST API works fine)"))
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", uh.GetAll)
		r.Post("/", uh.CreateRecord)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", uh.ReadRecord)
			r.Put("/", uh.UpdateRecord)
			r.Delete("/", uh.DeleteRecord)
		})
	})

	r.Route("/admins", func(r chi.Router) {
		r.Get("/", ah.GetAll)
		r.Post("/", ah.CreateRecord)

		//r.Route("/{id}", func(r chi.Router) {
		//	r.Get("/", uh.ReadRecord)
		//	r.Put("/", uh.UpdateRecord)
		//	r.Delete("/", uh.DeleteRecord)
		//})
	})

	return r
}
