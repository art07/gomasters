package entity

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
	"go.uber.org/zap"
	"time"
)

type Admin struct {
	ID        string    `validate:"required,uuid"`
	Firstname string    `validate:"required,alpha,min=3,max=20" json:"Firstname"`
	Lastname  string    `validate:"required,alpha,min=3,max=20" json:"Lastname"`
	Email     string    `validate:"required,email" json:"Email"`
	Age       int       `validate:"required,numeric,gte=0,lte=100" json:"Age"`
	Created   time.Time `validate:"required"`
}

func NewAdmin() *User {
	return &User{
		ID:      uuid.New().String(),
		Created: time.Now(),
	}
}

func (a *Admin) String() string {
	return fmt.Sprintf("Id > %v, first name > %s, last name > %s", a.ID, a.Firstname, a.Lastname)
}

func (a *Admin) Validate(l *zap.Logger) bool {
	validate := validator.New()
	if err := validate.Struct(a); err != nil {
		l.Error("Validation error for admin", zap.Error(err), zap.String("user", a.String()))
		return false
	}
	return true
}
