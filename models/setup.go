// package models

// import (
// 	"context"
// 	"log"
// 	"os"

// 	_ "github.com/jackc/pgx/v4/pgxpool"
// 	"github.com/joho/godotenv"
// )

// func SetupModels() (*pgxpool.Pool, error) {
// 	ctx := context.Background()
// 	err := godotenv.Load(`.env`)

// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	dsn := os.Getenv("DSN")

// 	db, err := pgxpool.Connect(ctx, dsn)
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}

// 	_, err = db.Exec(ctx, `
// 	CREATE TABLE IF NOT EXISTS  users (
//     	id BIGSERIAL PRIMARY KEY,
//     	name VARCHAR(150) NOT NULL,
//     	login VARCHAR(60) UNIQUE NOT NULL,
//     	password VARCHAR(255) NOT NULL,
// 		created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 	);
// 	CREATE TABLE IF NOT EXISTS  rooms (
// 		id BIGSERIAL PRIMARY KEY,
// 		name VARCHAR(150),
// 		first_member_id BIGINT REFERENCES users,
// 		second_member_id BIGINT REFERENCES users,
// 		created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 	);
	
// 	CREATE TABLE IF NOT EXISTS  messages (
// 		id BIGSERIAL PRIMARY KEY,
// 		room_id BIGINT REFERENCES rooms,
// 		created_by BIGINT REFERENCES users,
// 		text TEXT NOT NULL,
// 		created TIMESTAMP DEFAULT CURRENT_TIMESTAMP
// 	);
	
// 	CREATE TABLE IF NOT EXISTS tokens 
// 	(
// 		token TEXT NOT NULL UNIQUE,
// 		user_id BIGINT NOT NULL REFERENCES users,
// 		expire  TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP + INTERVAL '1 hour',
// 		created TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
// 	);
// 	`)

// 	if err != nil {
// 		log.Print(err)
// 		return nil, err
// 	}

// 	return db, nil
// 
