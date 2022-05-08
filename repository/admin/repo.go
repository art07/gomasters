package admin

import (
	"database/sql"
	"fmt"
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
	rows, err := ar.db.Query("SELECT * FROM admins;")
	if err != nil {
		return nil, fmt.Errorf("error in GetAll (admins) > %v", err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer rows.Close()

	var persons []entity.Person
	for rows.Next() {
		var a entity.Admin
		if err := rows.Scan(&a.ID, &a.Firstname, &a.Lastname, &a.Email, &a.Age, &a.Created); err != nil {
			ar.logger.Error("Admin reading error", zap.Error(err))
			continue
		}

		// Struct validation.
		if ok := a.Validate(ar.logger); !ok {
			continue
		}

		persons = append(persons, &a)
		ar.logger.Info("Admin added to slice", zap.String("user", a.String()))
	}
	return persons, nil
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
