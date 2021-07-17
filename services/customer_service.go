package services

import (
	"github.com/vu-nhan/gin-dig-grpc/repositories"
)

type customerService struct {
	customerRepository repositories.CustomerRepository
}

type CustomerService interface {
	GetAll() string
	GetByCode(code string) string
}

func NewCustomerService(injectedCustomerRepository repositories.CustomerRepository) CustomerService {
	return &customerService{
		customerRepository: injectedCustomerRepository,
	}
}

func (s *customerService) GetAll() string {
	return s.customerRepository.GetAll()
}

func (s *customerService) GetByCode(code string) string {
	return s.customerRepository.GetByCode(code)
}
