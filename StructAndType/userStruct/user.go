package user

import (
	"errors"
	"fmt"
	"time"
)

type User struct {
	FirstName string
	LastName  string
	BirthDate string
	CreatedAt time.Time
}
type Admin struct {
	Email    string
	Password string
	User
}

func (u *User) PrintUserDetails() {
	fmt.Println(u.FirstName, u.LastName, u.BirthDate)
}

func (u *User) ClearUserName() {
	u.FirstName = ""
	u.LastName = ""
}
func New(userfirstName string, userlastName string, userbirthdate string) (*User, error) {
	if userfirstName == "" || userlastName == "" || userbirthdate == "" {
		return nil, errors.New("First Name,Last Name, Birthday can't be empty")
	}
	return &User{
		FirstName: userfirstName,
		LastName:  userlastName,
		BirthDate: userbirthdate,
		CreatedAt: time.Now(),
	}, nil
}
func NewAdmin(email string, password string) *Admin {
	return &Admin{
		Email:    email,
		Password: password,
		User: User{
			FirstName: "Admin",
			LastName:  "Admin",
			BirthDate: "___",
			CreatedAt: time.Now(),
		},
	}
}
