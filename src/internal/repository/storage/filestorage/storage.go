package filestorage

import(
	"myApi/internal/domain/user"
	"log"
)

type StorageFile struct {
	File     *os.File
}

func New() *StorageFile{
	users := make(map[int] user.User)
	return &StorageMap{
		Users: users,
	}
}
 
func (s *StorageFile) CreateUser(newUser user.User) error{
	log.Println("Пользователь создан")
	s.Users[newUser.Id] = newUser
	return nil
}

func (s *StorageFile) ReadUser(id int) user.User{
	return s.Users[id]
}

func (s *StorageFile) UpdateUser(updatedUser user.User) {
	s.Users[updatedUser.Id] = updatedUser
}

func (s *StorageFile) DeleteUser(id int) {
	delete(s.Users, id)
}
