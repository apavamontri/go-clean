package interfaces

import "github.com/apavamontri/go-clean/domain"

type MemoryRepository struct {
}

func NewMemoryRepository() *MemoryRepository {
	return &MemoryRepository{}
}

func (m *MemoryRepository) GetEmployee(id int) (domain.Employee, error) {
	emp := domain.Employee{}
	emp.ID = 1
	emp.FirstName = "John"
	emp.LastName = "Doe"

	return emp, nil
}
