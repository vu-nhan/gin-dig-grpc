package handlers

import (
	"github.com/vu-nhan/gin-dig-grpc/services"
)

type CustomerHandler struct {
	customerService services.CustomerService
}

func NewCustomerHandler(injectedCustomerService services.CustomerService) CustomerHandler {
	return CustomerHandler{
		customerService: injectedCustomerService,
	}
}

func (h * CustomerHandler) GetAllVehicle() string {
	return h.customerService.GetAll()
}

func (h *CustomerHandler) GetVehicleByCode() string{
	return h.customerService.GetByCode("")
}