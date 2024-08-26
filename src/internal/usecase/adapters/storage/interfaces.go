package storage

import (
	"myApi/internal/domain/user"
)

//Данный интерфейс будет имплементирован в слое repository

type User interface{
	CreateEmployee(newUser *user.User) error
	GetAllEmployees() (*user.Employees, error) 
	// ReadUser(id int) user.User
	// UpdateUser(user user.User)
	// DeleteUser(id int) 
}