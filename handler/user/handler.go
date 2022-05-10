package user

import (
	"encoding/json"
	"fmt"
	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"go.uber.org/zap"
	"net/http"
	"playground/rest-api/gomasters/entity"
)

type Usecase interface {
	GetAll() ([]*entity.User, error)
	Create(*entity.User) (string, error)
	GetById(id string) (*entity.User, error)
	Update(string, *entity.User) (string, error)
	Delete(recordId string) (string, error)
}

type Handler struct {
	logger *zap.Logger
	uc     Usecase
}

func NewHandler(l *zap.Logger, uc Usecase) *Handler {
	return &Handler{
		logger: l, uc: uc,
	}
}

func (h *Handler) GetAll(w http.ResponseWriter, _ *http.Request) {
	users, err := h.uc.GetAll()
	if err != nil {
		h.logger.Error("get all error", zap.Error(err))
		render(w, "get all error")
		return
	}
	h.logger.Info("ger all succeeded")

	render(w, users)
}

func (h *Handler) Create(w http.ResponseWriter, r *http.Request) {
	//goland:noinspection GoUnhandledErrorResult
	defer r.Body.Close()

	u := entity.NewUser()
	if err := json.NewDecoder(r.Body).Decode(u); err != nil {
		h.logger.Error("decode user error", zap.Error(err))
		render(w, "decode user error")
		return
	}

	userId, err := h.uc.Create(u)
	if err != nil {
		h.logger.Error("create user error", zap.Error(err))
		render(w, "create user error")
		return
	}

	render(w, fmt.Sprintf("User with ID: %s, created successfully!", userId))
}

func (h *Handler) GetById(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := checkUUID(id); err != nil {
		h.logger.Error("uuid error", zap.Error(err))
		render(w, "uuid error")
		return
	}

	user, err := h.uc.GetById(id)
	if err != nil {
		h.logger.Error("get by id error", zap.Error(err))
		render(w, "get by id error")
		return
	}
	render(w, user)
}

func (h *Handler) Update(w http.ResponseWriter, r *http.Request) {
	//goland:noinspection GoUnhandledErrorResult
	defer r.Body.Close()

	id := mux.Vars(r)["id"]
	if err := checkUUID(id); err != nil {
		h.logger.Error("uuid error", zap.Error(err))
		render(w, "uuid error")
		return
	}

	user, err := h.uc.GetById(id)
	if err != nil {
		h.logger.Error("update error, user not found", zap.Error(err))
		render(w, "update error, user not found")
		return
	}

	if err = json.NewDecoder(r.Body).Decode(user); err != nil {
		h.logger.Error("decode user error", zap.Error(err))
		render(w, "decode user error")
		return
	}

	userId, err := h.uc.Update(id, user)
	if err != nil {
		h.logger.Error("update error", zap.Error(err))
		render(w, "update error")
		return
	}

	render(w, fmt.Sprintf("User with ID: %s, updated successfully!", userId))
}

func (h *Handler) Delete(w http.ResponseWriter, r *http.Request) {
	id := mux.Vars(r)["id"]
	if err := checkUUID(id); err != nil {
		h.logger.Error("uuid error, can't delete user", zap.Error(err))
		render(w, "uuid error, can't delete user")
		return
	}

	userId, err := h.uc.Delete(id)
	if err != nil {
		h.logger.Error("delete user error", zap.Error(err))
		render(w, "delete user error")
		return
	}

	render(w, fmt.Sprintf("User with ID: %s, deleted successfully!", userId))
}

func checkUUID(userId string) error {
	_, err := uuid.Parse(userId)
	return err
}

func render(w http.ResponseWriter, data interface{}) {
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(data)
}
