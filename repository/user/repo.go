package user

import (
	"database/sql"
	"fmt"
	"go.uber.org/zap"
	"playground/rest-api/gomasters/entity"
)

type UserRepository struct {
	logger *zap.Logger
	db     *sql.DB
}

func NewUserRepository(l *zap.Logger, db *sql.DB) *UserRepository {
	return &UserRepository{
		logger: l,
		db:     db,
	}
}

func (ur *UserRepository) GetAll() ([]entity.Person, error) {
	rows, err := ur.db.Query("SELECT * FROM users;")
	if err != nil {
		return nil, fmt.Errorf("error in GetAll (users) > %v", err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer rows.Close()

	var persons []entity.Person
	for rows.Next() {
		var u entity.User
		if err := rows.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Age, &u.Created); err != nil {
			ur.logger.Error("User reading error", zap.Error(err))
			continue
		}

		// Struct validation.
		if ok := u.Validate(ur.logger); !ok {
			continue
		}

		persons = append(persons, &u)
		ur.logger.Info("User added to slice", zap.String("user", u.String()))
	}
	return persons, nil
}

func (ur *UserRepository) CreateRecord(p entity.Person) (string, error) {
	u := p.(*entity.User)
	row := ur.db.QueryRow(
		"INSERT INTO users(id, first_name, last_name, email, age, created) "+
			"VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;",
		u.ID, u.Firstname, u.Lastname, u.Email, u.Age, u.Created)
	if row.Err() != nil {
		return "", fmt.Errorf("insert error > %v", row.Err())
	}

	var id string
	if err := row.Scan(&id); err != nil {
		return "", fmt.Errorf("insert error > %v", err)
	}

	return id, nil
}

func (ur *UserRepository) ReadRecord(id string) (entity.Person, error) {
	var u entity.User
	row := ur.db.QueryRow("SELECT * FROM users WHERE id=$1;", id)
	if err := row.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Age, &u.Created); err != nil {
		return nil, fmt.Errorf("read record error")
	}

	return &u, nil
}

func (ur *UserRepository) UpdateRecord(recordId string, p entity.Person) (string, error) {
	u := p.(*entity.User)
	row := ur.db.QueryRow(
		"UPDATE users SET id=$1, first_name=$2, last_name=$3, email=$4, age=$5, created=$6 WHERE id=$7 RETURNING id;",
		u.ID, u.Firstname, u.Lastname, u.Email, u.Age, u.Created, recordId)
	if row.Err() != nil {
		return "", fmt.Errorf("update error > %v", row.Err())
	}

	var id string
	if err := row.Scan(&id); err != nil {
		return "", fmt.Errorf("insert error > %v", err)
	}

	return id, nil
}

func (ur *UserRepository) DeleteRecord(recordId string) (string, error) {
	row := ur.db.QueryRow(
		"DELETE FROM users WHERE id=$1 RETURNING id;", recordId)
	if row.Err() != nil {
		return "", fmt.Errorf("delete error > %v", row.Err())
	}

	var id string
	if err := row.Scan(&id); err != nil {
		return "", fmt.Errorf("delete error > %v", err)
	}

	return id, nil
}
