package models

type User struct {
	ID         int64
	FirstName  string
	SecondName string
	Email      string
	Role       string // event owner / user

}
