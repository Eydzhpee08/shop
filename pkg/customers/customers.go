package customers

import (
	"context"
	"log"

	"github.com/jackc/pgx/v4"
)

type Customer struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Phone       string    `json:"phone"`
	Email       string	  `json:"email"`
}

type Service struct {
	db *pgx.Conn
}

func NewTeacherService(db *pgx.Conn) *Service {
	return &Service{db: db}
}

func (s *Service) GetCustomers(ctx context.Context, limit, offset int64) ([]Customer, error) {
	sql := `SELECT id, name, phone, email created FROM customers LIMIT $1 OFFSET $2`
	rows, err := s.db.Query(ctx, sql, limit, offset)
	if err != nil {
		return nil, err
	}
	customers := make([]Customer, 0)
	for rows.Next() {
		customer := Customer{}
		err = rows.Scan(
			&customer.ID,
			&customer.Name,
			&customer.Phone,
			&customer.Email,
		)
		if err != nil {
			return nil, err
		}
		customers = append(customers, customer)

	}
	return customers, nil
}

func (s *Service) AddCustomer(ctx context.Context, customer Customer) error {

	sql := "INSERT INTO customers (name, phone, email) VALUES ($1, $2, $3);"
	_, err := s.db.Exec(ctx, sql, customer.Name, customer.Phone, customer.Email)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

func (s *Service) EditCustomer(ctx context.Context, customer Customer) error {

	sql := "UPDATE customers SET name=$1, phone=$2,email=$3 WHERE id=$4;"
	exec, err := s.db.Exec(ctx, sql, customer.Name, customer.Phone, customer.ID, customer.Email)
	if err != nil {
		log.Print(err)
		return err
	}
	affected := exec.RowsAffected()
	if affected == 0 {
		return pgx.ErrNoRows
	}

	return nil
}

func (s *Service) RemoveCustomer(ctx context.Context, customerID int) error {

	sql := "DELETE FROM customers WHERE id=$1;"
	exec, err := s.db.Exec(ctx, sql, customerID)
	if err != nil {
		log.Print(err)
		return err
	}
	if affected := exec.RowsAffected(); affected == 0 {
		return pgx.ErrNoRows
	}
	return nil
}
