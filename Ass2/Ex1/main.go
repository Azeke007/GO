package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

// Database connection string
const (
	connStr = "host=localhost port=5432 user=postgres password=AK_qwerty dbname=GOlang sslmode=disable"
)
// Connect to PostgreSQL database
func ConnectToDB() (*sql.DB, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	fmt.Println("Successfully connected to PostgreSQL database")
	return db, nil
}
// Create users table
func CreateTable(db *sql.DB) error {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT NOT NULL,
		age INT NOT NULL
	);`

	_, err := db.Exec(query)
	if err != nil {
		return err
	}
	fmt.Println("Table users created or already exists.")
	return nil
}
// Insert a new user into the table
func InsertUser(db *sql.DB, name string, age int) error {
	query := `INSERT INTO users (name, age) VALUES ($1, $2);`
	_, err := db.Exec(query, name, age)
	if err != nil {
		return err
	}
	fmt.Printf("User %s (age %d) inserted successfully.\n", name, age)
	return nil
}
// Query all users from the table
func QueryUsers(db *sql.DB) error {
	query := `SELECT id, name, age FROM users;`
	rows, err := db.Query(query)
	if err != nil {
		return err
	}
	defer rows.Close()

	fmt.Println("Users:")
	for rows.Next() {
		var id int
		var name string
		var age int
		err = rows.Scan(&id, &name, &age)
		if err != nil {
			return err
		}
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", id, name, age)
	}

	return rows.Err()
} 

func main() {
	//  Connect to the database
	db, err := ConnectToDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	//  Create the table
	err = CreateTable(db)
	if err != nil {
		log.Fatal(err)
	}
	//   Insert a new user
	err = InsertUser(db, "Azeke", 22)
	if err != nil {
		log.Fatal(err)
	}
	err = InsertUser(db, "Qora", 23)
	if err != nil {
		log.Fatal(err)
	}
	//   Query all users from the table
	err = QueryUsers(db)
	if err != nil {
		log.Fatal(err)
	}
}
