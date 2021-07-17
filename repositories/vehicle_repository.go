package repositories

import "fmt"

type vehicleRepository struct {
}

type VehicleRepository interface {
	GetAll() string
	GetByCode(code string) string
}

func NewVehicleRepository() VehicleRepository {
	return &vehicleRepository{}
}

func (s *vehicleRepository) GetAll() string {
	fmt.Println("Get All Vehicle")
	return "Get All Vehicle"
}

func (s *vehicleRepository) GetByCode(code string) string {
	fmt.Println("Get Vehicle By Code")
	return "Get Vehicle By Code"
}
