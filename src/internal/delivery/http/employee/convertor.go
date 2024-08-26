package employee

import (
	"myApi/internal/domain/user"
	"time"
	"log"
)

//конвертация для структуры сотрудниов для перехода из слоёв Domain->Delivery
func ToEmployeesDelivery(employees *user.Employees) (EmployeesDelivery) {
	size := employees.Size()
	emplsDel := NewEmployees(size)
	for _, employee := range employees.Data {
		emplDel := ToEmployeeDelivery(employee)
		emplsDel.Append(emplDel)
	}
	return emplsDel
}

//конвертация для структуры сотрудника для перехода из слоёв Domain->Delivery
func ToEmployeeDelivery(employee user.User) EmployeeDelivery {
	strBirthDay := employee.BirthDay.Format("2006-01-02")
	return EmployeeDelivery {
		Name : employee.Name,
		SurName :employee.SurName,
		BirthDay : strBirthDay,
	}
}

//конвертация для структуры сотрудника для перехода из слоёв Domain->Delivery
func ToEmployee(employeeReq EmployeeDelivery) (*user.User, error) {
	birthDay, err := time.Parse("2006-01-02", employeeReq.BirthDay)
	if err != nil {
        log.Println(err)
		return nil, err
    }
	result := &user.User{
		Name :  	employeeReq.Name,
		SurName: 	employeeReq.SurName,
		BirthDay: 	birthDay,
	}
	return result, err 
}

// func ToEmployee(*employee user.User) (*EmployeeDelivery, error) {
// 	birthDay, err := time.Parse("2006-01-02", employeeReq.BirthDay)
// 	if err != nil {
//         log.Println(err)
// 		return nil, err
//     }
// 	result := &user.User{
// 		Name :  	employeeReq.Name,
// 		SurName: 	employeeReq.SurName,
// 		BirthDay: 	birthDay,
// 	}
// 	return result, err 
// }
