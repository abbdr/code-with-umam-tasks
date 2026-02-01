package database

import (
	"database/sql"
	"log"

	_ "github.com/lib/pq"
)

func InitDB(connectionString string) (*sql.DB, error) {
	// Open database
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}

	// Test connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}

	// Set connection pool settings (optional tapi recommended)
	db.SetMaxOpenConns(25)
	db.SetMaxIdleConns(5)

	log.Println("Database connected successfully")
	return db, nil

	// conn, err := pgx.Connect(context.Background(), connectionString)
	// if err != nil {
	// 	log.Fatalf("Failed to connect to the database: %v", err)
	// }
	// defer conn.Close(context.Background())

	// // Example query to test connection
	// var version string
	// if err := conn.QueryRow(context.Background(), "SELECT version()").Scan(&version); err != nil {
	// 	log.Fatalf("Query failed: %v", err)
	// }

	// log.Println("Connected to:", version)

	// return conn, nil
}
