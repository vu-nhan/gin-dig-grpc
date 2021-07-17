package handlers

import (
	"github.com/vu-nhan/gin-dig-grpc/services"
)

type VehicleHandler struct {
	vehicleService services.VehicleService
}

func NewVehicleHandler(injectedVehicleService services.VehicleService) VehicleHandler {
	return VehicleHandler{
		vehicleService: injectedVehicleService,
	}
}

func (h * VehicleHandler) GetAllVehicle() string {
	return h.vehicleService.GetAll()
}

func (h *VehicleHandler) GetVehicleByCode() string{

	return h.vehicleService.GetByCode("")
}