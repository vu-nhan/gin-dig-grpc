package services

import (
	"github.com/vu-nhan/gin-dig-grpc/repositories"
)

type vehicleService struct {
	vehicleRepository repositories.VehicleRepository
}

type VehicleService interface {
	GetAll() string
	GetByCode(code string) string
}

func NewVehicleService(injectedVehicleRepository repositories.VehicleRepository) VehicleService {
	return &vehicleService{
		vehicleRepository: injectedVehicleRepository,
	}
}

func (s *vehicleService) GetAll() string {
	return s.vehicleRepository.GetAll()
}

func (s *vehicleService) GetByCode(code string) string {
	return s.vehicleRepository.GetByCode(code)
}
