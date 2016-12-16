package main

import "github.com/apavamontri/go-clean/usecases"
import "github.com/apavamontri/go-clean/interfaces"
import "net/http"

func main() {
	employeeInteractor := usecases.NewEmployeeInteractor()
	employeeInteractor.EmployeeRepository = interfaces.NewMemoryRepository()

	webHandler := interfaces.NewWebHandler()
	webHandler.EmployeeInteractor = employeeInteractor

	http.HandleFunc("/", webHandler.GetEmployee)
	http.ListenAndServe(":3000", nil)
}
