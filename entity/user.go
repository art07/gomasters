package entity

import (
	"fmt"
	"github.com/google/uuid"
	"time"
)

type User struct {
	ID        string    `validate:"required,uuid"`
	Firstname string    `validate:"required,alpha,min=3,max=20" json:"Firstname"`
	Lastname  string    `validate:"required,alpha,min=3,max=20" json:"Lastname"`
	Email     string    `validate:"required,email" json:"Email"`
	Age       int       `validate:"required,numeric,gte=0,lte=100" json:"Age"`
	Created   time.Time `validate:"required"`
}

func NewUser() *User {
	return &User{
		ID:      uuid.New().String(),
		Created: time.Now(),
	}
}

func (u *User) String() string {
	return fmt.Sprintf("Id > %v, first name > %s, last name > %s", u.ID, u.Firstname, u.Lastname)
}
