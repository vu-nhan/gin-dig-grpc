package services

import (
	"github.com/vu-nhan/gin-dig-grpc/repositories"
)

type customerService struct {
	customerRepository repositories.CustomerRepository
}

type CustomerService interface {
	GetAll() (string, error)
	GetByCode(code string) (string, error)
}

func NewCustomerService(injectedCustomerRepository repositories.CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: injectedCustomerRepository,
	}
}

func (s *customerService) GetAll() (string, error) {
	return s.customerRepository.GetAll(), nil
}

func (s *customerService) GetByCode(code string) (string, error) {
	return s.customerRepository.GetByCode(code), nil
}
