package usecases

import "github.com/apavamontri/go-clean/domain"

type EmployeeRepository interface {
	GetEmployee(id int) (domain.Employee, error)
}

type EmployeeInteractor struct {
	EmployeeRepository
}

func NewEmployeeInteractor() *EmployeeInteractor {
	return &EmployeeInteractor{}
}
