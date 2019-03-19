package models

import (
	"fmt"
	"math/rand"
)

type User struct {
	ID        int64
	FirstName string
	LastName  string
	salary    int32
}

func NewUser(name string, surname string, salary int32) *User {
	return &User{
		ID:        rand.Int63(),
		FirstName: name,
		LastName:  surname,
		salary:    salary,
	}
}

func (u *User) GetSalary() {
	fmt.Println(u.FirstName, u.salary+3000)
}

func (u *User) GetPrintNameUser() {
	//panic("implement me")
	fmt.Println(u.FirstName)
}
