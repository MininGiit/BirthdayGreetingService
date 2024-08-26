package mapstorage

import(
	"myApi/internal/domain/user"
	"log"
)

type StorageMap struct {
	Users map[int] user.User
}

func New() *StorageMap{
	users := make(map[int] user.User)
	return &StorageMap{
		Users: users,
	}
}
 
func (s *StorageMap) CreateUser(newUser user.User) error{
	log.Println("Пользователь создан")
	s.Users[newUser.Id] = newUser
	return nil
}

func (s *StorageMap) ReadUser(id int) user.User{
	return s.Users[id]
}

func (s *StorageMap) UpdateUser(updatedUser user.User) {
	s.Users[updatedUser.Id] = updatedUser
}

func (s *StorageMap) DeleteUser(id int) {
	delete(s.Users, id)
}
