package repositories

import "fmt"

type customerRepository struct {

}

type CustomerRepository interface {
	GetAll() string
	GetByCode(code string) string
}

func NewCustomerRepository() CustomerRepository {
	return &customerRepository{}
}


func (s *customerRepository) GetAll() string {
	fmt.Println("Get All Customers")
	return "Get All Customers"
}

func (s *customerRepository) GetByCode(code string) string {
	fmt.Println("Get Customer By Code")
	return "Get Customer By Code"
}