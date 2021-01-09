package bill

import (
	"context"
	"log"
	"github.com/jackc/pgx/v4"
)

type Bill struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
}

type Service struct {
	db *pgx.Conn
}

func NewBillService(db *pgx.Conn) *Service {
	return &Service{db: db}
}

func (s *Service) GetBill(ctx context.Context, limit, offset int64) ([]Bill, error) {
	sql := `select id, name, created from bill limit $1 offset $2`
	rows, err := s.db.Query(ctx, sql, limit, offset)
	if err != nil {
		return nil, err
	}
	bills := make([]Bill, 0)
	for rows.Next() {
		bill := Bill{}
		err = rows.Scan(
			&bill.ID,
			&bill.Name,
		)
		if err != nil {
			return nil, err
		}
		bills = append(bills, bill)

	}
	return bills, nil
}

func (s *Service) AddBill(ctx context.Context, bill Bill) error {

	sql := "INSERT INTO bills (name) VALUES ($1);"
	_, err := s.db.Exec(ctx, sql, bill.Name)
	if err != nil {
		log.Print(err)
		return err
	}

	return nil
}

