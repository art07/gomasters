package user

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"playground/rest-api/gomasters/entity"
)

type Repository interface {
	GetAll() ([]*entity.User, error)
	Create(*entity.User) (string, error)
	GetById(id string) (*entity.User, error)
	Update(string, *entity.User) (string, error)
	Delete(recordId string) (string, error)
}

type Usecase struct {
	repo Repository
}

func NewUsecase(r Repository) *Usecase {
	return &Usecase{
		repo: r,
	}
}

func (u *Usecase) GetAll() ([]*entity.User, error) {
	return u.repo.GetAll()
}

func (u *Usecase) Create(user *entity.User) (string, error) {
	if err := validate(user); err != nil {
		return "", fmt.Errorf("validation error: %v", err)
	}

	return u.repo.Create(user)
}

func (u *Usecase) GetById(userId string) (*entity.User, error) {
	return u.repo.GetById(userId)
}

func (u *Usecase) Update(userId string, user *entity.User) (string, error) {
	if err := validate(user); err != nil {
		return "", fmt.Errorf("validation error: %v", err)
	}

	return u.repo.Update(userId, user)
}

func (u *Usecase) Delete(userId string) (string, error) {
	return u.repo.Delete(userId)
}

func validate(user *entity.User) error {
	v := validator.New()
	return v.Struct(user)
}
