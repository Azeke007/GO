package main

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

const (
	host     = "localhost"
	port     = 5432
	user     = "postgres"
	password = "AK_qwerty"
	dbname   = "GOlang"
)
// Инициализация пула соединений с PostgreSQL
func connectDB() (*sql.DB, error) {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		host, port, user, password, dbname)

	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		return nil, err
	}

	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	fmt.Println("Successfully connected to PostgreSQL")
	return db, nil
}
// Создание таблицы с ограничениями
func createTable(db *sql.DB) {
	query := `
	CREATE TABLE IF NOT EXISTS users (
		id SERIAL PRIMARY KEY,
		name TEXT UNIQUE NOT NULL,
		age INT NOT NULL
	)`
	_, err := db.Exec(query)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Table created successfully")
}
// Вставка данных с транзакцией
func insertUsersWithTransaction(db *sql.DB, users []User) error {
	tx, err := db.Begin()
	if err != nil {
		return err
	}

	for _, user := range users {
		_, err = tx.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", user.Name, user.Age)
		if err != nil {
			tx.Rollback()
			return err
		}
	}

	err = tx.Commit()
	if err != nil {
		return err
	}

	fmt.Println("Users inserted successfully")
	return nil
}
// Запрос данных с фильтрацией и пагинацией
func queryUsersWithPagination(db *sql.DB, minAge int, limit int, offset int) ([]User, error) {
	query := `
	SELECT id, name, age FROM users WHERE age >= $1 ORDER BY id LIMIT $2 OFFSET $3
	`
	rows, err := db.Query(query, minAge, limit, offset)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.Name, &user.Age)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	return users, nil
}
// Обновление данных
func updateUser(db *sql.DB, id int, name string, age int) error {
	query := `UPDATE users SET name = $1, age = $2 WHERE id = $3`
	_, err := db.Exec(query, name, age, id)
	if err != nil {
		return err
	}
	fmt.Printf("User with ID %d updated successfully\n", id)
	return nil
}
// Удаление данных
func deleteUser(db *sql.DB, id int) error {
	query := `DELETE FROM users WHERE id = $1`
	_, err := db.Exec(query, id)
	if err != nil {
		return err
	}
	fmt.Printf("User with ID %d deleted successfully\n", id)
	return nil
}

type User struct {
	ID   int
	Name string
	Age  int
}

func main() {
	db, err := connectDB()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	createTable(db)

	users := []User{
		{Name: "Alice", Age: 30},
		{Name: "Bob", Age: 25},
		{Name: "Charlie", Age: 35},
	}

	err = insertUsersWithTransaction(db, users)
	if err != nil {
		log.Fatalf("Error inserting users: %v", err)
	}

	minAge := 20
	limit := 2
	offset := 0
	result, err := queryUsersWithPagination(db, minAge, limit, offset)
	if err != nil {
		log.Fatalf("Error querying users: %v", err)
	}

	fmt.Println("Queried Users:")
	for _, user := range result {
		fmt.Printf("ID: %d, Name: %s, Age: %d\n", user.ID, user.Name, user.Age)
	}

	err = updateUser(db, 1, "Alice Updated", 31)
	if err != nil {
		log.Fatalf("Error updating user: %v", err)
	}

	err = deleteUser(db, 2)
	if err != nil {
		log.Fatalf("Error deleting user: %v", err)
	}
}
