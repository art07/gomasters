package user

import (
	"errors"
	"github.com/stretchr/testify/assert"
	"playground/rest-api/gomasters/entity"
	"playground/rest-api/gomasters/mock"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
)

func TestUsecase_GetAll(t *testing.T) {
	type expected struct {
		Users []*entity.User
		Err   error
	}

	type payload struct {
		GetMockRepo func(*gomock.Controller, []*entity.User, error) *mock.MockRepository
	}

	tc := []struct {
		name     string
		expected expected
		payload  payload
	}{
		{
			name: "get all users success",
			expected: expected{
				Users: []*entity.User{
					{
						ID:        "1636c7ff-e1bc-40d1-a368-3cdbfc2dd97c",
						Firstname: "SecondUser",
						Lastname:  "LastNameB",
						Email:     "user2@gmail.com",
						Age:       21,
						Created:   time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        "f2a44f36-0956-4019-9134-bbb0a2f63b01",
						Firstname: "ThirdUser",
						Lastname:  "LastNameC",
						Email:     "user3@gmail.com",
						Age:       22,
						Created:   time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC),
					},
					{
						ID:        "1d2ef152-f440-4be2-b659-46cc6dcbc966",
						Firstname: "FirstUser",
						Lastname:  "LastNameA",
						Email:     "user1@gmail.com",
						Age:       20,
						Created:   time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC),
					},
				},
				Err: nil,
			},
			payload: payload{
				GetMockRepo: func(mockCtrl *gomock.Controller, users []*entity.User, err error) *mock.MockRepository {
					mockRepo := mock.NewMockRepository(mockCtrl)
					mockRepo.EXPECT().GetAll().Return(users, err).Times(1)
					return mockRepo
				}},
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := test.payload.GetMockRepo(mockCtrl, test.expected.Users, test.expected.Err)
			usecase := NewUsecase(mockRepo)
			users, err := usecase.GetAll()

			assert.Nil(t, err)
			assert.ElementsMatch(t, users, test.expected.Users)
		})
	}
}

func TestUsecase_Create(t *testing.T) {
	type expected struct {
		Id  string
		Err error
	}

	type payload struct {
		User        *entity.User
		GetMockRepo func(*gomock.Controller, *entity.User, string, error) *mock.MockRepository
	}

	tc := []struct {
		name     string
		expected expected
		payload  payload
	}{
		{
			name: "create success",
			expected: expected{
				Id:  "1d2ef152-f440-4be2-b659-46cc6dcbc966",
				Err: nil,
			},
			payload: payload{
				User: &entity.User{
					ID:        "1d2ef152-f440-4be2-b659-46cc6dcbc966",
					Firstname: "FirstUser",
					Lastname:  "LastNameA",
					Email:     "user1@gmail.com",
					Age:       20,
					Created:   time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC),
				},
				GetMockRepo: func(mockCtrl *gomock.Controller, user *entity.User, id string, err error) *mock.MockRepository {
					mockRepo := mock.NewMockRepository(mockCtrl)
					mockRepo.EXPECT().Create(user).Return(id, err).Times(1)
					return mockRepo
				}},
		},
		{
			name: "validation error",
			expected: expected{
				Id:  "",
				Err: errors.New("validation error: Key: 'User.Firstname' Error:Field validation for 'Firstname' failed on the 'alpha' tag"),
			},
			payload: payload{
				User: &entity.User{
					ID:        "1d2ef152-f440-4be2-b659-46cc6dcbc966",
					Firstname: "FirstUser100",
					Lastname:  "LastNameA",
					Email:     "user1@gmail.com",
					Age:       20,
					Created:   time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC),
				},
				GetMockRepo: func(mockCtrl *gomock.Controller, user *entity.User, id string, err error) *mock.MockRepository {
					return mock.NewMockRepository(mockCtrl)
				}},
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := test.payload.GetMockRepo(mockCtrl, test.payload.User, test.expected.Id, test.expected.Err)
			usecase := NewUsecase(mockRepo)
			resId, err := usecase.Create(test.payload.User)

			if err != nil {
				assert.EqualError(t, err, test.expected.Err.Error())
				assert.NotNil(t, err)
				assert.EqualValues(t, test.expected.Id, resId)
				return
			}

			assert.Nil(t, err)
			assert.EqualValues(t, test.expected.Id, resId)
		})
	}
}

func TestUsecase_GetById(t *testing.T) {
	type expected struct {
		User *entity.User
		Err  error
	}

	type payload struct {
		UserId      string
		GetMockRepo func(*gomock.Controller, string, *entity.User, error) *mock.MockRepository
	}

	tc := []struct {
		name     string
		expected expected
		payload  payload
	}{
		{
			name: "get user success",
			expected: expected{
				User: &entity.User{
					ID:        "1636c7ff-e1bc-40d1-a368-3cdbfc2dd97c",
					Firstname: "SecondUser",
					Lastname:  "LastNameB",
					Email:     "user2@gmail.com",
					Age:       21,
					Created:   time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC),
				},
				Err: nil,
			},
			payload: payload{
				UserId: "1636c7ff-e1bc-40d1-a368-3cdbfc2dd97c",
				GetMockRepo: func(mockCtrl *gomock.Controller, userId string, user *entity.User, err error) *mock.MockRepository {
					mockRepo := mock.NewMockRepository(mockCtrl)
					mockRepo.EXPECT().GetById(userId).Return(user, err).Times(1)
					return mockRepo
				}},
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := test.payload.GetMockRepo(mockCtrl, test.payload.UserId, test.expected.User, test.expected.Err)
			usecase := NewUsecase(mockRepo)
			user, err := usecase.GetById(test.payload.UserId)

			assert.Nil(t, err)
			assert.EqualValues(t, test.expected.User, user)
		})
	}
}

func TestUsecase_Update(t *testing.T) {
	type expected struct {
		Id  string
		Err error
	}

	type payload struct {
		UserId      string
		User        *entity.User
		GetMockRepo func(*gomock.Controller, string, *entity.User, string, error) *mock.MockRepository
	}

	tc := []struct {
		name     string
		expected expected
		payload  payload
	}{
		{
			name: "update success",
			expected: expected{
				Id:  "1d2ef152-f440-4be2-b659-46cc6dcbc966",
				Err: nil,
			},
			payload: payload{
				User: &entity.User{
					ID:        "1d2ef152-f440-4be2-b659-46cc6dcbc966",
					Firstname: "FirstUser",
					Lastname:  "LastNameA",
					Email:     "user1@gmail.com",
					Age:       20,
					Created:   time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC),
				},
				GetMockRepo: func(mockCtrl *gomock.Controller, userId string, user *entity.User, id string, err error) *mock.MockRepository {
					mockRepo := mock.NewMockRepository(mockCtrl)
					mockRepo.EXPECT().Update(userId, user).Return(id, err).Times(1)
					return mockRepo
				}},
		},
		{
			name: "validation error",
			expected: expected{
				Id:  "",
				Err: errors.New("validation error: Key: 'User.Firstname' Error:Field validation for 'Firstname' failed on the 'alpha' tag"),
			},
			payload: payload{
				User: &entity.User{
					ID:        "1d2ef152-f440-4be2-b659-46cc6dcbc966",
					Firstname: "FirstUser100",
					Lastname:  "LastNameA",
					Email:     "user1@gmail.com",
					Age:       20,
					Created:   time.Date(2022, time.Month(5), 7, 0, 0, 0, 0, time.UTC),
				},
				GetMockRepo: func(mockCtrl *gomock.Controller, userId string, user *entity.User, id string, err error) *mock.MockRepository {
					return mock.NewMockRepository(mockCtrl)
				}},
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := test.payload.GetMockRepo(mockCtrl, test.payload.UserId, test.payload.User, test.expected.Id, test.expected.Err)
			usecase := NewUsecase(mockRepo)
			resId, err := usecase.Update(test.payload.UserId, test.payload.User)

			if err != nil {
				assert.EqualError(t, err, test.expected.Err.Error())
				assert.NotNil(t, err)
				assert.EqualValues(t, test.expected.Id, resId)
				return
			}

			assert.Nil(t, err)
			assert.EqualValues(t, test.expected.Id, resId)
		})
	}
}

func TestUsecase_Delete(t *testing.T) {
	type expected struct {
		UserId string
		Err    error
	}

	type payload struct {
		UserId      string
		GetMockRepo func(*gomock.Controller, string, string, error) *mock.MockRepository
	}

	tc := []struct {
		name     string
		expected expected
		payload  payload
	}{
		{
			name: "delete user success",
			expected: expected{
				UserId: "1636c7ff-e1bc-40d1-a368-3cdbfc2dd97c",
				Err:    nil,
			},
			payload: payload{
				UserId: "1636c7ff-e1bc-40d1-a368-3cdbfc2dd97c",
				GetMockRepo: func(mockCtrl *gomock.Controller, userIdIn string, userIdOut string, err error) *mock.MockRepository {
					mockRepo := mock.NewMockRepository(mockCtrl)
					mockRepo.EXPECT().Delete(userIdIn).Return(userIdOut, err).Times(1)
					return mockRepo
				}},
		},
	}

	for _, test := range tc {
		t.Run(test.name, func(t *testing.T) {
			mockCtrl := gomock.NewController(t)
			defer mockCtrl.Finish()

			mockRepo := test.payload.GetMockRepo(mockCtrl, test.payload.UserId, test.expected.UserId, test.expected.Err)
			usecase := NewUsecase(mockRepo)
			userId, err := usecase.Delete(test.payload.UserId)

			assert.Nil(t, err)
			assert.EqualValues(t, test.expected.UserId, userId)
		})
	}
}
