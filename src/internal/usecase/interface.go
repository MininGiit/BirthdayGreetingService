package usecase

import (
	"myApi/internal/domain/user"
)
//пока не понимаю зачем???
type Employees interface{
	Create(user *user.User) error
	GetAll() (*user.Employees, error)
	Subscribe(name, surName) error
	UnSubscribe(name, surName) error
 	// Read(id int) 
	// Update(user user.User)
	// Delete(id int) 
}
