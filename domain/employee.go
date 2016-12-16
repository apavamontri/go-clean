package domain

import (
	"fmt"
)

type EmployeeRepository interface {
	GetEmployee(id int) (Employee, error)
}

type Employee struct {
	ID        int
	FirstName string
	LastName  string
}

func (e *Employee) String() string {
	return fmt.Sprintf("%s %s", e.FirstName, e.LastName)
}
