package router

import (
	gochi "github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
)

type Handler interface {
	GetAll(w http.ResponseWriter, r *http.Request)
	CreateRecord(w http.ResponseWriter, r *http.Request)
	ReadRecord(w http.ResponseWriter, r *http.Request)
	UpdateRecord(w http.ResponseWriter, r *http.Request)
	DeleteRecord(w http.ResponseWriter, r *http.Request)
}

func NewRouter(uh, ah Handler, l *zap.Logger) http.Handler {
	r := gochi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("REST API works fine)")); err != nil {
			l.Error("Write index page error", zap.Error(err))
		}
	})

	r.Route("/users", func(r gochi.Router) {
		r.Get("/", uh.GetAll)
		r.Post("/", uh.CreateRecord)

		r.Route("/{id}", func(r gochi.Router) {
			r.Get("/", uh.ReadRecord)
			r.Put("/", uh.UpdateRecord)
			r.Delete("/", uh.DeleteRecord)
		})
	})

	r.Route("/admins", func(r gochi.Router) {
		r.Get("/", ah.GetAll)
		r.Post("/", ah.CreateRecord)

		r.Route("/{id}", func(r gochi.Router) {
			r.Get("/", ah.ReadRecord)
			r.Put("/", ah.UpdateRecord)
			r.Delete("/", ah.DeleteRecord)
		})
	})

	return r
}
