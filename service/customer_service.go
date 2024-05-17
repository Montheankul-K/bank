package service

import (
	"database/sql"
	"errors"
	"github.com/Montheankul-K/bank/repository"
	"log"
)

type customerService struct {
	customerRepository repository.CustomerRepository // refer to port
}

func NewCustomerService(customerRepository repository.CustomerRepository) customerService {
	return customerService{
		customerRepository: customerRepository,
	}
}

func (s customerService) GetCustomers() ([]CustomerResponse, error) {
	customers, err := s.customerRepository.GetAll()
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customers not found")
		}
		log.Println(err)

		return nil, err
	}

	customerResponses := []CustomerResponse{}
	for _, customer := range customers {
		customerResponse := CustomerResponse{
			CustomerID: customer.CustomerID,
			Name:       customer.Name,
			Status:     customer.Status,
		}
		customerResponses = append(customerResponses, customerResponse)
	}

	return customerResponses, nil
}

func (s customerService) GetCustomer(id int) (*CustomerResponse, error) {
	customer, err := s.customerRepository.GetByID(id)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, errors.New("customer not found")
		}
		log.Println(err)

		return nil, err
	}

	customerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &customerResponse, nil
}
