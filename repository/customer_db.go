package repository

import (
	"github.com/jmoiron/sqlx"
)

// adapter
type customerRepositoryDB struct {
	db *sqlx.DB
}

// constructor : we should not exposed adapter so we need to use constructor
func NewCustomerRepositoryDB(db *sqlx.DB) customerRepositoryDB {
	return customerRepositoryDB{
		db: db,
	}
}

// conform (implement) interface
func (r customerRepositoryDB) GetAll() ([]Customer, error) {
	customers := []Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers"
	err := r.db.Select(&customers, query)
	if err != nil {
		// repository don't need to handle error or log because it not business logic
		return nil, err
	}

	return customers, nil
}

func (r customerRepositoryDB) GetByID(id int) (*Customer, error) {
	customer := Customer{}
	query := "select customer_id, name, date_of_birth, city, zipcode, status from customers where customer_id = ?"
	err := r.db.Get(&customer, query, id)
	if err != nil {
		return nil, err
	}

	return &customer, nil
}
