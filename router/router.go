package router

import (
	"database/sql"
	"github.com/go-chi/chi/v5"
	"go.uber.org/zap"
	"net/http"
	userHandler "playground/rest-api/gomasters/handler/user"
	userRepo "playground/rest-api/gomasters/repository/postgres/user"
	userUsecase "playground/rest-api/gomasters/usecase/user"
)

func NewRouter(db *sql.DB, l *zap.Logger) chi.Router {
	// DB inject in repository
	uRepo := userRepo.NewRepository(db)
	//aRepo := adminRepo.NewRepository(db)

	// Repo inject in usecase
	uUsecase := userUsecase.NewUsecase(uRepo)
	//aUsecase := adminUsecase.NewUsecase(aRepo)

	// Usecase inject in handler
	uHandler := userHandler.NewHandler(l, uUsecase)
	//aHandler := adminHandler.NewHandler(l, aUsecase)

	r := chi.NewRouter()

	r.Get("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("REST API works fine)")); err != nil {
			l.Error("Write index page error", zap.Error(err))
		}
	})

	r.Route("/users", func(r chi.Router) {
		r.Get("/", uHandler.GetAll)
		r.Post("/", uHandler.Create)

		r.Route("/{id}", func(r chi.Router) {
			r.Get("/", uHandler.GetById)
			r.Put("/", uHandler.Update)
			r.Delete("/", uHandler.Delete)
		})
	})

	//r.Route("/admins", func(r chi.Router) {
	//	r.Get("/", aHandler.GetAll)
	//	r.Post("/", aHandler.Create)
	//
	//	r.Route("/{id}", func(r chi.Router) {
	//		r.Get("/", aHandler.GetById)
	//		r.Put("/", aHandler.Update)
	//		r.Delete("/", aHandler.Delete)
	//	})
	//})

	return r
}
