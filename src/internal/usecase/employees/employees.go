package employees

import (
	"myApi/internal/domain/user"
)

func (uc *UseCase) Create(user *user.User) error{
	return uc.adapterStorage.CreateEmployee(user)
}

func (uc *UseCase) GetAll() (*user.Employees, error) {
	return uc.adapterStorage.GetAllEmployees()
}