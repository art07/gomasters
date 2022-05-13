package user

import (
	"database/sql"
	"errors"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/stretchr/testify/assert"
	"playground/rest-api/gomasters/entity"
	"testing"
	"time"
)

const dbString = "user=postgres password=postgres host=localhost port=5432 database=gomasters-db-test sslmode=disable"

func TestRepository_Create(t *testing.T) {
	type expected struct {
		UserId string
		Err    error
	}

	type payload struct {
		User *entity.User
	}

	tc := []struct {
		name     string
		expected expected
		payload  payload
	}{
		{
			name: "create success",
			expected: expected{
				UserId: "a01c6ae7-86c1-400e-beb2-5a5c6e15785c",
				Err:    nil,
			},
			payload: payload{
				User: &entity.User{
					ID:        "a01c6ae7-86c1-400e-beb2-5a5c6e15785c",
					Firstname: "NewUser",
					Lastname:  "NewUserLastname",
					Email:     "newuser@gmail.com",
					Age:       30,
					Created:   time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC),
				},
			},
		},
		{
			name: "insert error",
			expected: expected{
				UserId: "",
				Err:    errors.New("create error: ERROR: value too long for type character varying(40) (SQLSTATE 22001)"),
			},
			payload: payload{
				User: &entity.User{
					ID:        "a01c6ae7-86c1-400e-beb2-5a5c6e15785c",
					Firstname: "NewUserrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrrr",
					Lastname:  "NewUserLastname",
					Email:     "newuser@gmail.com",
					Age:       30,
					Created:   time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC),
				},
			},
		},
	}

	db, _ := sql.Open("pgx", dbString)
	//goland:noinspection GoUnhandledErrorResult
	defer db.Close()

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			userRepo := NewRepository(db)
			userId, err := userRepo.Create(test.payload.User)

			if err != nil {
				assert.EqualError(t, err, test.expected.Err.Error())
				assert.EqualValues(t, test.expected.UserId, userId)
				return
			}

			assert.Nil(t, err)
			assert.EqualValues(t, test.expected.UserId, userId)
		})
	}
}

func TestRepository_GetAll(t *testing.T) {
	type expected struct {
		Users []*entity.User
		Err   error
	}

	type payload struct {
		GetPostgres func() *sql.DB
	}

	tc := []struct {
		name     string
		expected expected
		payload  payload
	}{
		{
			name: "db error",
			expected: expected{
				Users: nil,
				Err:   errors.New("get all users query error: cannot parse `some db string`: failed to parse as DSN (invalid dsn)"),
			},
			payload: payload{
				GetPostgres: func() *sql.DB {
					db, _ := sql.Open("pgx", "some db string")
					return db
				},
			},
		},
		{
			name: "get all success",
			expected: expected{
				Users: []*entity.User{
					{
						ID:        "a01c6ae7-86c1-400e-beb2-5a5c6e15785c",
						Firstname: "NewUser",
						Lastname:  "NewUserLastname",
						Email:     "newuser@gmail.com",
						Age:       30,
						Created:   time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC),
					},
				},
				Err: nil,
			},
			payload: payload{
				GetPostgres: func() *sql.DB {
					db, _ := sql.Open("pgx", dbString)
					return db
				},
			},
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			db := test.payload.GetPostgres()
			//goland:noinspection GoUnhandledErrorResult
			defer db.Close()

			userRepo := NewRepository(db)
			users, err := userRepo.GetAll()

			if err != nil {
				assert.EqualError(t, err, test.expected.Err.Error())
				assert.Nil(t, users)
				return
			}

			assert.Nil(t, err)
			assert.ElementsMatch(t, test.expected.Users, users)
		})
	}
}

func TestRepository_GetById(t *testing.T) {
	type expected struct {
		User *entity.User
		Err  error
	}

	type payload struct {
		UserId      string
		GetPostgres func() *sql.DB
	}

	tc := []struct {
		name     string
		expected expected
		payload  payload
	}{
		{
			name: "db error",
			expected: expected{
				User: nil,
				Err:  errors.New("get user by id error: cannot parse `some db string`: failed to parse as DSN (invalid dsn)"),
			},
			payload: payload{
				UserId: "a01c6ae7-86c1-400e-beb2-5a5c6e15785c",
				GetPostgres: func() *sql.DB {
					db, _ := sql.Open("pgx", "some db string")
					return db
				},
			},
		},
		{
			name: "get user success",
			expected: expected{
				User: &entity.User{
					ID:        "a01c6ae7-86c1-400e-beb2-5a5c6e15785c",
					Firstname: "NewUser",
					Lastname:  "NewUserLastname",
					Email:     "newuser@gmail.com",
					Age:       30,
					Created:   time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC),
				},
				Err: nil,
			},
			payload: payload{
				UserId: "a01c6ae7-86c1-400e-beb2-5a5c6e15785c",
				GetPostgres: func() *sql.DB {
					db, _ := sql.Open("pgx", dbString)
					return db
				},
			},
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			db := test.payload.GetPostgres()
			//goland:noinspection GoUnhandledErrorResult
			defer db.Close()

			userRepo := NewRepository(db)
			user, err := userRepo.GetById(test.payload.UserId)

			if err != nil {
				assert.EqualError(t, err, test.expected.Err.Error())
				assert.Nil(t, user)
				return
			}

			assert.Nil(t, err)
			assert.EqualValues(t, test.expected.User, user)
		})
	}
}

func TestRepository_Update(t *testing.T) {
	type expected struct {
		UserIdOut string
		Err       error
	}

	type payload struct {
		UserIdIn    string
		User        *entity.User
		GetPostgres func() *sql.DB
	}

	tc := []struct {
		name     string
		expected expected
		payload  payload
	}{
		{
			name: "db error",
			expected: expected{
				UserIdOut: "",
				Err:       errors.New("update error: cannot parse `some db string`: failed to parse as DSN (invalid dsn)"),
			},
			payload: payload{
				UserIdIn: "a01c6ae7-86c1-400e-beb2-5a5c6e15785c",
				User: &entity.User{
					ID:        "a01c6ae7-86c1-400e-beb2-5a5c6e15785c",
					Firstname: "NewUserUPD",
					Lastname:  "NewUserLastnameUPD",
					Email:     "newuser@gmail.com",
					Age:       30,
					Created:   time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC),
				},
				GetPostgres: func() *sql.DB {
					db, _ := sql.Open("pgx", "some db string")
					return db
				},
			},
		},
		{
			name: "update success",
			expected: expected{
				UserIdOut: "a01c6ae7-86c1-400e-beb2-5a5c6e15785c",
				Err:       nil,
			},
			payload: payload{
				UserIdIn: "a01c6ae7-86c1-400e-beb2-5a5c6e15785c",
				User: &entity.User{
					ID:        "a01c6ae7-86c1-400e-beb2-5a5c6e15785c",
					Firstname: "NewUserUPD",
					Lastname:  "NewUserLastnameUPD",
					Email:     "newuser@gmail.com",
					Age:       30,
					Created:   time.Date(2022, time.Month(7), 7, 0, 0, 0, 0, time.UTC),
				},
				GetPostgres: func() *sql.DB {
					db, _ := sql.Open("pgx", dbString)
					return db
				},
			},
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			db := test.payload.GetPostgres()
			//goland:noinspection GoUnhandledErrorResult
			defer db.Close()

			userRepo := NewRepository(db)
			userId, err := userRepo.Update(test.payload.UserIdIn, test.payload.User)

			if err != nil {
				assert.EqualError(t, err, test.expected.Err.Error())
				assert.EqualValues(t, test.expected.UserIdOut, userId)
				return
			}

			assert.Nil(t, err)
			assert.EqualValues(t, test.expected.UserIdOut, userId)
		})
	}
}

func TestRepository_Delete(t *testing.T) {
	type expected struct {
		UserId string
		Err    error
	}

	type payload struct {
		UserId      string
		GetPostgres func() *sql.DB
	}

	tc := []struct {
		name     string
		expected expected
		payload  payload
	}{
		{
			name: "delete error / db error",
			expected: expected{
				UserId: "",
				Err:    errors.New("delete error: cannot parse `wrong db string`: failed to parse as DSN (invalid dsn)"),
			},
			payload: payload{
				UserId: "a01c6ae7-86c1-400e-beb2-5a5c6e15785a",
				GetPostgres: func() *sql.DB {
					db, _ := sql.Open("pgx", "wrong db string")
					return db
				},
			},
		},
		{
			name: "delete error / rows affected error",
			expected: expected{
				UserId: "",
				Err:    errors.New("no row found to delete"),
			},
			payload: payload{
				UserId: "a01c6ae7-86c1-400e-beb2-5a5c6e15785a",
				GetPostgres: func() *sql.DB {
					db, _ := sql.Open("pgx", dbString)
					return db
				},
			},
		},
		{
			name: "delete success",
			expected: expected{
				UserId: "a01c6ae7-86c1-400e-beb2-5a5c6e15785c",
				Err:    nil,
			},
			payload: payload{
				UserId: "a01c6ae7-86c1-400e-beb2-5a5c6e15785c",
				GetPostgres: func() *sql.DB {
					db, _ := sql.Open("pgx", dbString)
					return db
				},
			},
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			db := test.payload.GetPostgres()
			//goland:noinspection GoUnhandledErrorResult
			defer db.Close()

			userRepo := NewRepository(db)
			userId, err := userRepo.Delete(test.payload.UserId)

			if err != nil {
				assert.EqualError(t, err, test.expected.Err.Error())
				assert.EqualValues(t, test.expected.UserId, userId)
				return
			}

			assert.Nil(t, err)
			assert.EqualValues(t, test.expected.UserId, userId)
		})
	}
}
