package router

import (
	"database/sql"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	userHandler "playground/rest-api/gomasters/handler/user"
	userRepo "playground/rest-api/gomasters/repository/postgres/user"
	userUsecase "playground/rest-api/gomasters/usecase/user"
)

func NewRouter(db *sql.DB, l *zap.Logger) *mux.Router {
	// DB inject in repository
	uRepo := userRepo.NewRepository(db)
	//aRepo := adminRepo.NewRepository(db)

	// Repo inject in usecase
	uUsecase := userUsecase.NewUsecase(uRepo)
	//aUsecase := adminUsecase.NewUsecase(aRepo)

	// Usecase inject in handler
	uHandler := userHandler.NewHandler(l, uUsecase)
	//aHandler := adminHandler.NewHandler(l, aUsecase)

	r := mux.NewRouter()

	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if _, err := w.Write([]byte("REST API works fine)")); err != nil {
			l.Error("Write index page error", zap.Error(err))
		}
	})

	usersRouter := r.PathPrefix("/users").Subrouter()
	usersRouter.HandleFunc("", uHandler.GetAll).Methods(http.MethodGet)
	usersRouter.HandleFunc("", uHandler.Create).Methods(http.MethodPost)

	usersIdRouter := usersRouter.PathPrefix("/{id}").Subrouter()
	usersIdRouter.HandleFunc("", uHandler.GetById).Methods(http.MethodGet)
	usersIdRouter.HandleFunc("", uHandler.Update).Methods(http.MethodPut)
	usersIdRouter.HandleFunc("", uHandler.Delete).Methods(http.MethodDelete)

	return r
}
