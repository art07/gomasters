package user

import (
	"database/sql"
	"errors"
	"fmt"
	"playground/rest-api/gomasters/entity"
)

type Repository struct {
	db *sql.DB
}

func NewRepository(db *sql.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (ur *Repository) GetAll() ([]*entity.User, error) {
	rows, err := ur.db.Query("SELECT * FROM users;")
	if err != nil {
		return nil, fmt.Errorf("get all users query error: %v", err)
	}
	//goland:noinspection GoUnhandledErrorResult
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var u entity.User
		if err = rows.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Age, &u.Created); err != nil {
			return nil, fmt.Errorf("get all users rows scan error: %v", err)
		}

		users = append(users, &u)
	}

	return users, nil
}

func (ur *Repository) Create(u *entity.User) (string, error) {
	row := ur.db.QueryRow(
		"INSERT INTO users(id, first_name, last_name, email, age, created) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id;",
		u.ID, u.Firstname, u.Lastname, u.Email, u.Age, u.Created)
	if row.Err() != nil {
		return "", fmt.Errorf("create error: %v", row.Err())
	}

	var userId string
	if err := row.Scan(&userId); err != nil {
		return "", fmt.Errorf("scan id of created user error: %v", err)
	}

	return userId, nil
}

func (ur *Repository) GetById(id string) (*entity.User, error) {
	var u entity.User
	row := ur.db.QueryRow("SELECT * FROM users WHERE id=$1;", id)
	if row.Err() != nil {
		return nil, fmt.Errorf("get user by id error: %v", row.Err())
	}

	if err := row.Scan(&u.ID, &u.Firstname, &u.Lastname, &u.Email, &u.Age, &u.Created); err != nil {
		return nil, fmt.Errorf("get user by id row scan error: %v", err)
	}

	return &u, nil
}

func (ur *Repository) Update(userId string, u *entity.User) (string, error) {
	row := ur.db.QueryRow(
		"UPDATE users SET id=$1, first_name=$2, last_name=$3, email=$4, age=$5, created=$6 WHERE id=$7 RETURNING id;",
		u.ID, u.Firstname, u.Lastname, u.Email, u.Age, u.Created, userId)
	if row.Err() != nil {
		return "", fmt.Errorf("update error: %v", row.Err())
	}

	var id string
	if err := row.Scan(&id); err != nil {
		return "", fmt.Errorf("update ok but row scan for id error: %v", err)
	}

	return id, nil
}

func (ur *Repository) Delete(userId string) (string, error) {
	res, err := ur.db.Exec("DELETE FROM users WHERE id=$1;", userId)
	if err != nil {
		return "", fmt.Errorf("delete error: %v", err)
	}

	rowsAffected, _ := res.RowsAffected()
	if rowsAffected != 1 {
		return "", errors.New("no row found to delete")
	}

	return userId, nil
}
