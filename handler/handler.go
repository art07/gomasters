package handler

import "playground/rest-api/gomasters/entity"

type Repository interface {
	GetAll() ([]entity.Person, error)
	CreateRecord(entity.Person) (string, error)
	ReadRecord(string) (entity.Person, error)
	UpdateRecord(string, entity.Person) (string, error)
	DeleteRecord(string) (string, error)
}
