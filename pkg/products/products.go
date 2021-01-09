package products

import (
	"context"
	"log"
	"github.com/jackc/pgx/v4"
)

//Products ...
type Product struct {
	ID          int64     `json:"id"`
	Name        string    `json:"name"`
	Price     int64     `json:"price"`
}

type Service struct {
	db *pgx.Conn
}

func NewProductService(db *pgx.Conn) *Service {
	return &Service{db: db}
}


func (s *Service) GetProduct(ctx context.Context, limit, offset int64) ([]Product, error) {
	sql := `SELECT id, name, price, created FROM products LIMIT $1 OFFSET $2`
	rows, err := s.db.Query(ctx, sql, limit, offset)
	if err != nil {
		return nil, err
	}
	products := make([]Product, 0)
	for rows.Next() {
		product := Product{}
		err = rows.Scan(
			&product.ID,
			&product.Name,
			&product.Price,
		)
		if err != nil {
			return nil, err
		}
		products = append(products, product)

	}
	return products, nil
}

func (s *Service) AddProducts(ctx context.Context, product Product) error {
	if product.ID == 0 {
		sql := "INSERT INTO product (name, price) VALUES ($1, $2);"
		_, err := s.db.Exec(ctx, sql, product.Name, product.Price)
		if err != nil {
			log.Print(err)
			return err
		}
	}

	return nil
}

func (s *Service) EditProducts(ctx context.Context, product Product) error {
	if product.ID != 0 {
		sql := "UPDATE products SET name=$1, price=$2 WHERE id=$3;"
		_, err := s.db.Exec(ctx, sql, product.Name, product.Price, product.ID)
		if err != nil {
			log.Print(err)
			return err
		}
	}

	return nil
}

func (s *Service) RemoveProducts(ctx context.Context, productID int) error {
	
	sql := "DELETE FROM products WHERE id=$1;"
	exec, err := s.db.Exec(ctx, sql, productID)
	if err != nil {
		log.Print(err)
		return err
	}

	if affected:=exec.RowsAffected(); affected==0{
		return pgx.ErrNoRows
	}
	return nil
}