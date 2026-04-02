package config

import (
	"context"
	"fmt"
	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnectDB() {
	var err error
	// Replace with your PostgreSQL credentials
	connStr := "postgres://task_user:task_pass@localhost:5432/task_db?sslmode=disable"

	DB, err = pgxpool.New(context.Background(), connStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Connected to PostgreSQL")
}