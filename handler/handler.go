package handler

import (
	"encoding/json"
	"fmt"
	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"playground/rest-api/gomasters/entity"
)

type Repository interface {
	GetAll() ([]entity.Person, error)
	CreateRecord(entity.Person) (string, error)
	ReadRecord(string) (entity.Person, error)
	UpdateRecord(string, entity.Person) (string, error)
	DeleteRecord(string) (string, error)
}

type Handler struct {
	logger *zap.Logger
	repo   Repository
}

func NewHandler(l *zap.Logger, r Repository) *Handler {
	return &Handler{logger: l, repo: r}
}

func (h *Handler) GetAll(w http.ResponseWriter, r *http.Request) {
	res, err := h.repo.GetAll()
	if err != nil {
		h.logger.Error(err.Error())
		render("Get all records error (see logs for more info)", w)
		return
	}
	render(res, w)
}

func (h *Handler) CreateRecord(w http.ResponseWriter, r *http.Request) {
	//goland:noinspection GoUnhandledErrorResult
	defer r.Body.Close()

	u := entity.NewUser()
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		h.logger.Error(err.Error())
		render("Create record error (see logs for more info)", w)
		return
	}

	// Struct validation.
	if ok := u.Validate(h.logger); !ok {
		render("Create record error (see logs for more info)", w)
		return
	}

	res, err := h.repo.CreateRecord(u)
	if err != nil {
		h.logger.Error(err.Error())
		render("Create record error (see logs for more info)", w)
		return
	}

	render(fmt.Sprintf("Record with ID > %s created successfully!", res), w)
}

func (h *Handler) ReadRecord(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	_, err := uuid.Parse(id)
	if err != nil {
		h.logger.Error(err.Error())
		render("Read record error (see logs for more info)", w)
		return
	}

	res, err := h.repo.ReadRecord(id)
	if err != nil {
		h.logger.Error(err.Error())
		render("Read record error (see logs for more info)", w)
		return
	}
	render(res, w)
}

func (h *Handler) UpdateRecord(w http.ResponseWriter, r *http.Request) {
	//goland:noinspection GoUnhandledErrorResult
	defer r.Body.Close()

	id := chi.URLParam(r, "id")

	// UUID check
	_, err := uuid.Parse(id)
	if err != nil {
		h.logger.Error(err.Error())
		render("Update record by ID error (see logs for more info)", w)
		return
	}

	// Entity check
	if _, err := h.repo.ReadRecord(id); err != nil {
		h.logger.Error("Record not found", zap.String("userId", id), zap.Error(err))
		render("Record not found (see logs for more info)", w)
		return
	}

	// Decode data for updating
	var u entity.User
	if err := json.NewDecoder(r.Body).Decode(&u); err != nil {
		h.logger.Error(err.Error())
		render("Decode record error (see logs for more info)", w)
		return
	}

	// Update section
	resId, err := h.repo.UpdateRecord(id, &u)
	if err != nil {
		h.logger.Error("Update error", zap.Error(err))
		render("Update error (see logs for more info)", w)
		return
	}

	render(fmt.Sprintf("Record with ID > %s updated successfully!", resId), w)
}

func (h *Handler) DeleteRecord(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// UUID check
	_, err := uuid.Parse(id)
	if err != nil {
		h.logger.Error(err.Error())
		render("Delete by ID error (see logs for more info)", w)
		return
	}

	res, err := h.repo.DeleteRecord(id)
	if err != nil {
		h.logger.Error(err.Error())
		render("Delete record error (see logs for more info)", w)
		return
	}

	render(fmt.Sprintf("Record with ID > %s deleted successfully!", res), w)
}

func render(data interface{}, w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
	//goland:noinspection GoUnhandledErrorResult
	json.NewEncoder(w).Encode(data)
}
