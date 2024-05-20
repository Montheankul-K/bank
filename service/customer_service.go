package service

import (
	"database/sql"
	"errors"
	"github.com/Montheankul-K/bank/errs"
	"github.com/Montheankul-K/bank/logs"
	"github.com/Montheankul-K/bank/repository"
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
		if errors.Is(err, sql.ErrNoRows) {
			return nil, errs.NewNotFoundError("customer not found")
		}

		logs.Error(err)
		return nil, errs.NewUnexpectedError()
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
		if errors.Is(err, sql.ErrNoRows) {
			/*
				return nil, errs.AppError{
					Code:    http.StatusNotFound,
					Message: "customer not found",
				}
			*/

			return nil, errs.NewNotFoundError("customer not found")
		}

		logs.Error(err)
		/*
			return nil, errs.AppError{
				Code:    http.StatusInternalServerError,
				Message: "unexpected error",
			}
		*/

		return nil, errs.NewUnexpectedError()
	}

	customerResponse := CustomerResponse{
		CustomerID: customer.CustomerID,
		Name:       customer.Name,
		Status:     customer.Status,
	}

	return &customerResponse, nil
}
