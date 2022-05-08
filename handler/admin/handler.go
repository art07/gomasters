package admin

import (
	"encoding/json"
	"fmt"
	gochi "github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"net/http"
	"playground/rest-api/gomasters/entity"
	"playground/rest-api/gomasters/handler"
)

//goland:noinspection GoNameStartsWithPackageName
type AdminHandler struct {
	logger *zap.Logger
	repo   handler.Repository
}

func NewAdminHandler(l *zap.Logger, r handler.Repository) *AdminHandler {
	return &AdminHandler{logger: l, repo: r}
}

func (h *AdminHandler) GetAll(w http.ResponseWriter, _ *http.Request) {
	res, err := h.repo.GetAll()
	if err != nil {
		h.logger.Error(err.Error())
		render("Get all records error (see logs for more info)", w)
		return
	}
	render(res, w)
}

func (h *AdminHandler) CreateRecord(w http.ResponseWriter, r *http.Request) {
	//goland:noinspection GoUnhandledErrorResult
	defer r.Body.Close()

	a := entity.NewAdmin()
	if err := json.NewDecoder(r.Body).Decode(a); err != nil {
		h.logger.Error(err.Error())
		render("Create record error (see logs for more info)", w)
		return
	}

	// Struct validation.
	if ok := a.Validate(h.logger); !ok {
		render("Create record error (see logs for more info)", w)
		return
	}

	res, err := h.repo.CreateRecord(a)
	if err != nil {
		h.logger.Error(err.Error())
		render("Create record error (see logs for more info)", w)
		return
	}

	render(fmt.Sprintf("Record with ID > %s created successfully!", res), w)
}

func (h *AdminHandler) ReadRecord(w http.ResponseWriter, r *http.Request) {
	id := gochi.URLParam(r, "id")
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

func (h *AdminHandler) UpdateRecord(w http.ResponseWriter, r *http.Request) {
	//goland:noinspection GoUnhandledErrorResult
	defer r.Body.Close()

	id := gochi.URLParam(r, "id")

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
	var a entity.Admin
	if err := json.NewDecoder(r.Body).Decode(&a); err != nil {
		h.logger.Error(err.Error())
		render("Decode record error (see logs for more info)", w)
		return
	}

	// Update section
	resId, err := h.repo.UpdateRecord(id, &a)
	if err != nil {
		h.logger.Error("Update error", zap.Error(err))
		render("Update error (see logs for more info)", w)
		return
	}

	render(fmt.Sprintf("Record with ID > %s updated successfully!", resId), w)
}

func (h *AdminHandler) DeleteRecord(w http.ResponseWriter, r *http.Request) {
	id := gochi.URLParam(r, "id")

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
	_ = json.NewEncoder(w).Encode(data)
}
