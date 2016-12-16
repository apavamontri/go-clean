package usecases

import "github.com/apavamontri/go-clean/domain"

type EmployeeInteractor struct {
	EmployeeRepository domain.EmployeeRepository
}

func NewEmployeeInteractor() *EmployeeInteractor {
	return &EmployeeInteractor{}
}

func (ei *EmployeeInteractor) GetEmployee(id int) (domain.Employee, error) {
	return ei.EmployeeRepository.GetEmployee(id)
}
