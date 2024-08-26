package http

import (
	"myApi/internal/usecase"
	"github.com/gorilla/mux"
	"net/http"
)

type Delivery struct {
	ucEmployees usecase.Employees
	router *mux.Router  
}

func New(useCaseEmployes usecase.Employees) *Delivery {
	router := mux.NewRouter()
	return &Delivery{
		ucEmployees: useCaseEmployes,
		router: router,
	}
}

func (d *Delivery) StartServer() {
	d.router.HandleFunc("/signin", Signin).Methods("POST")

	api := d.router.PathPrefix("/auth").Subrouter()
    api.Use(JwtMiddleware)
	api.HandleFunc("/employee", d.PostUser).Methods("POST")
	api.HandleFunc("/employee", d.GetUsers).Methods("GET")
	http.ListenAndServe(":8080", d.router)
}