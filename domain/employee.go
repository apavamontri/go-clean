package domain

import (
	"fmt"
)

type Employee struct {
	ID        int
	FirstName string
	LastName  string
}

func (e *Employee) String() string {
	return fmt.Sprintf("%d: %s %s", e.ID, e.FirstName, e.LastName)
}
