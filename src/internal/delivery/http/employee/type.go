package employee

//структура сотрудника для слоя Delivery
type EmployeeDelivery struct {
	Name string `json:"name"`
	SurName string `json:"sur_name"`
	BirthDay string `json:"birth_day"`
}

//структура сотрудников для слоя Delivery
type EmployeesDelivery struct {
	Data []EmployeeDelivery `json:"employees"`
}

//создание структуры пользователей
func NewEmployees(n int) EmployeesDelivery{
	data := make([]EmployeeDelivery, 0, n)
	return EmployeesDelivery{
		Data: data,
	}
}

func (empls *EmployeesDelivery) Append(employee EmployeeDelivery) {
	empls.Data = append(empls.Data, employee) 
}

