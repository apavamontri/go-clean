package interfaces

import (
	"net/http"

	"github.com/apavamontri/go-clean/domain"
)

type EmployeeInteractor interface {
	GetEmployee(id int) (domain.Employee, error)
}

type WebHandler struct {
	EmployeeInteractor
}

func NewWebHandler() *WebHandler {
	return &WebHandler{}
}

func (wh *WebHandler) GetEmployee(w http.ResponseWriter, r *http.Request) {
	emp, _ := wh.EmployeeInteractor.GetEmployee(1)

	w.Write([]byte(emp.String()))
}
