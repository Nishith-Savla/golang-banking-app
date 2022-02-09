package service

import (
	"github.com/Nishith-Savla/golang-banking-app/domain"
	"github.com/Nishith-Savla/golang-banking-app/dto"
	"github.com/Nishith-Savla/golang-banking-lib/errs"
)

//go:generate mockgen -destination=../mocks/service/mockCustomerService.go -package=service github.com/Nishith-Savla/golang-banking-app/service CustomerService
type CustomerService interface {
	GetAllCustomers(string) ([]dto.CustomerResponse, *errs.AppError)
	GetCustomer(string) (*dto.CustomerResponse, *errs.AppError)
}

type DefaultCustomerService struct {
	repo domain.CustomerRepository
}

func (s DefaultCustomerService) GetAllCustomers(status string) ([]dto.CustomerResponse, *errs.AppError) {
	if status == "active" {
		status = "1"
	} else if status == "inactive" {
		status = "0"
	} else {
		status = ""
	}
	customers, err := s.repo.FindAll(status)
	if err != nil {
		return nil, err
	}
	customerResponses := make([]dto.CustomerResponse, len(customers))
	for i, customer := range customers {
		customerResponses[i] = customer.ToDto()
	}
	return customerResponses, err
}

func (s DefaultCustomerService) GetCustomer(id string) (*dto.CustomerResponse, *errs.AppError) {
	c, err := s.repo.ById(id)
	if err != nil {
		return nil, err
	}

	response := c.ToDto()

	return &response, nil
}

func NewCustomerService(repository domain.CustomerRepository) DefaultCustomerService {
	return DefaultCustomerService{repository}
}
