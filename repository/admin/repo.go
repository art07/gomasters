package admin

import (
	"database/sql"
	"go.uber.org/zap"
	"playground/rest-api/gomasters/entity"
)

type AdminRepository struct {
	logger *zap.Logger
	db     *sql.DB
}

func NewAdminRepository(l *zap.Logger, db *sql.DB) *AdminRepository {
	return &AdminRepository{
		logger: l,
		db:     db,
	}
}

func (ar *AdminRepository) GetAll() ([]entity.Person, error) {
	return nil, nil
}

func (ar *AdminRepository) CreateRecord(p entity.Person) (string, error) {
	return "", nil
}

func (ar *AdminRepository) ReadRecord(id string) (entity.Person, error) {
	return nil, nil
}

func (ar *AdminRepository) UpdateRecord(id string, p entity.Person) (string, error) {
	return "", nil
}

func (ar *AdminRepository) DeleteRecord(recordId string) (string, error) {
	return "", nil
}
