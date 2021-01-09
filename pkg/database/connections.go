package database

import (
	"context"

	"github.com/jackc/pgx/v4"
)



func NewDBConnection(dsn string) *pgx.Conn{
	DB, err := pgx.Connect(context.Background(), dsn)
	if err != nil {
		panic(err)
	}
	
	return DB
}

func Init(db *pgx.Conn) error{
	// TOOD: init tables
	_, err:=db.Exec(context.Background(),
	`
	CREATE TABLE if not exists bill (
		product_name TEXT REFERENCES products
		customers_name TEXT REFERENCES customers
	);`)
	if err!=nil{
		return err
	}
	_, err=db.Exec(context.Background(),
	`CREATE TABLE if not exists customers (
		id BIGSERIAL PRIMARY KEY,
        name TEXT NOT NULL,
		tel TEXT NOT NULL UNIQUE,
	);`)
	if err!=nil{
		return err
	}
	_, err=db.Exec(context.Background(),
	`
	CREATE TABLE if not exists products (
		id BIGSERIAL PRIMARY KEY,
        price INTEGER NOT NULL CHECK (price >0),
        name TEXT NOT NULL,
		);
	`)
	if err!=nil{
		return err
	}

	return nil
}

