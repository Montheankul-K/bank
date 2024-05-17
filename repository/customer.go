package repository

// port
type Customer struct {
	CustomerID  int    `db:"customer_id"`
	Name        string `db:"name"`
	DateOfBirth string `db:"date_of_birth"`
	City        string `db:"city"`
	ZipCode     string `db:"zip_code"`
	Status      int    `db:"status"`
}

type CustomerRepository interface {
	GetAll() ([]Customer, error)
	GetByID(int) (*Customer, error) // struct can't return nil but *struct can
}
