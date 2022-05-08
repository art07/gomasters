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
	a := p.(*entity.Admin)
	row := ar.db.QueryRow(
		"INSERT INTO admins(id, first_name, last_name, email, age, created) "+
			"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;",
		a.ID, a.Firstname, a.Lastname, a.Email, a.Age, a.Created)
	if row.Err() != nil {
		return "", fmt.Errorf("insert error > %v", row.Err())
	}

	var id string
	if err := row.Scan(&id); err != nil {
		return "", fmt.Errorf("insert error > %v", err)
	}

	return id, nil
}

func (ar *AdminRepository) ReadRecord(id string) (entity.Person, error) {
	var a entity.Admin
	row := ar.db.QueryRow("SELECT * FROM admins WHERE id=$1;", id)
	if err := row.Scan(&a.ID, &a.Firstname, &a.Lastname, &a.Email, &a.Age, &a.Created); err != nil {
		return nil, fmt.Errorf("read record error")
	}

	return &a, nil
}

func (ar *AdminRepository) UpdateRecord(recordId string, p entity.Person) (string, error) {
	u := p.(*entity.Admin)
	row := ar.db.QueryRow(
		"UPDATE admins SET id=$1, first_name=$2, last_name=$3, email=$4, age=$5, created=$6 WHERE id=$7 RETURNING id;",
		u.ID, u.Firstname, u.Lastname, u.Email, u.Age, u.Created, recordId)
	if row.Err() != nil {
		return "", fmt.Errorf("update error > %v", row.Err())
	}

	var id string
	if err := row.Scan(&id); err != nil {
		return "", fmt.Errorf("update error > %v", err)
	}

	return id, nil
}

func (ar *AdminRepository) DeleteRecord(recordId string) (string, error) {
	return "", nil
}
