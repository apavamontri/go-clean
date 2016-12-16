package main

import "github.com/apavamontri/go-clean/usecases"
import "github.com/apavamontri/go-clean/interfaces"
import "net/http"

func main() {
	employeeInteractor := usecases.EmployeeInteractor{}
	employeeInteractor.EmployeeRepository = interfaces.MemoryRepository{}

	webHandler := interfaces.WebHandler{}
	webHandler.EmployeeInteractor = employeeInteractor

	http.HandleFunc("/", webHandler.GetEmployee)
	http.ListenAndServe(":3000", nil)
}
