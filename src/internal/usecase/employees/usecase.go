package employees

import(
	"myApi/internal/usecase/adapters/storage"
)

type UseCase struct{
	adapterStorage storage.User
	// --- поле хранит интерфейс, который имплементирует методы обращения к БД
}

//в качестве аргумента интерфейс, то есть можно положить любую структуру, 
// имплиментирующую этот интерфейс 
func New(storage storage.User) *UseCase {
	return &UseCase{
		adapterStorage: storage,
	}
}

//func GetBirthdayPeople()