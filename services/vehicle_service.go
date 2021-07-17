package services

import (
	"github.com/vu-nhan/gin-dig-grpc/repositories"
)

type vehicleService struct {
	vehicleRepository repositories.VehicleRepository
}

type VehicleService interface {
	GetAll() (string, error)
	GetByCode(code string) (string, error)
}

func NewVehicleService(injectedVehicleRepository repositories.VehicleRepository) VehicleService {
	return &vehicleService{
		vehicleRepository: injectedVehicleRepository,
	}
}

func (s *vehicleService) GetAll() (string, error) {
	return s.vehicleRepository.GetAll(), nil
}

func (s *vehicleService) GetByCode(code string) (string, error) {
	return s.vehicleRepository.GetByCode(code), nil
}
