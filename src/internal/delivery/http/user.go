package http

import (
	"fmt"
	"net/http"
	employeeDelivery "myApi/internal/delivery/http/employee"
	"encoding/json"
	"log"

)

func (d *Delivery) GetUsers(w http.ResponseWriter, r *http.Request) {
	username := r.Context().Value("username").(string)
	fmt.Fprintln(w, username)
	employees, err := d.ucEmployees.GetAll()
	if err != nil {
		log.Println(err)
		return
	}
	//fmt.Println(employees)
	var empsDel employeeDelivery.EmployeesDelivery
	empsDel = employeeDelivery.ToEmployeesDelivery(employees)

	fmt.Printf("%+v\n", empsDel)
	jsonData, err := json.Marshal(empsDel)

	jsonStr := string(jsonData)
	fmt.Println(jsonStr)
    if err != nil {
        log.Println("POsosososo")
        return
    }
	fmt.Fprintln(w, jsonStr)
}

func (d *Delivery) PostUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Method POST")
	var employeeDel employeeDelivery.EmployeeDelivery

	err := json.NewDecoder(r.Body).Decode(&employeeDel)
	if err != nil {
		log.Println("Error JSON decode")
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	employee, _:= employeeDelivery.ToEmployee(employeeDel)
	
	if err := d.ucEmployees.Create(employee); err != nil {
		log.Println(err)
	}
}

func (d *Delivery) DefaultUser(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "default")
}