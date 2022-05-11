package user

import (
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
		GetMockRepo func(controller *gomock.Controller, users []*entity.User, err error) *mock.MockRepository
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
