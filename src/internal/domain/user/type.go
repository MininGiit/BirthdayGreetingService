//Доменный слой, содержит сущности
package user

import("time")

type User struct{
	Id 		int
	Name	string
	SurName string
	BirthDay time.Time
	Subscription []int //ID сотрудников, на который подписан пользователь
}

type Employees struct{
	Data map[int] User
}

func (emps *Employees) Size() int{
	return len(emps.Data)
}

func (emps *Employees) AddEmployee(employ User){
	emps.Data[employ.Id] = employ
}

func New(id int, name string, surName string, birthDay time.Time) (*User) {
	return &User {
		Id:		 id,
		Name: 	 name,
		SurName: surName,
		BirthDay: birthDay,
	}
}

func NewEmployees() (*Employees) {
	data := make(map[int] User)
	return &Employees {
		Data : data,
	}
}
