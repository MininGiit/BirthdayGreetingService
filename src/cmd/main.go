package main

import(
	"myApi/internal/repository/storage/postgres"
	ucEmployees "myApi/internal/usecase/employees"
	deliveryHttp "myApi/internal/delivery/http"
)

func main() {
	repo, _:= postgres.New()
	ucEmployes := ucEmployees.New(repo)
	delivery := deliveryHttp.New(ucEmployes)
	delivery.StartServer()
}