package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type User struct {
	ID   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
	Age  int    `json:"age"`
}

// Переменные для хранения подключений
var dbSQL *sql.DB
var dbGORM *gorm.DB

// Подключение к базе данных через database/sql
func connectSQL() *sql.DB {
	connStr := "user=postgres password=AK_qwerty dbname=GOlang sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	return db
}

// Подключение к базе данных через GORM
func connectGORM() *gorm.DB {
	dsn := "host=localhost user=postgres password=AK_qwerty dbname=GOlang port=5432 sslmode=disable"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect to the database:", err)
	}
	db.AutoMigrate(&User{})
	return db
}

// Получение всех пользователей через SQL с фильтрацией и сортировкой
func getUsersSQL(w http.ResponseWriter, r *http.Request) {
	query := "SELECT id, name, age FROM users"
	queryParams := []interface{}{}
	queryFilters := ""
	querySort := ""

	ageFilter := r.URL.Query().Get("age")
	if ageFilter != "" {
		queryFilters += " WHERE age >= $1"
		queryParams = append(queryParams, ageFilter)
	}

	sortParam := r.URL.Query().Get("sort")
	if sortParam == "name" {
		querySort = " ORDER BY name ASC"
	}

	query = query + queryFilters + querySort

	rows, err := dbSQL.Query(query, queryParams...)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	defer rows.Close()

	var users []User
	for rows.Next() {
		var user User
		if err := rows.Scan(&user.ID, &user.Name, &user.Age); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		users = append(users, user)
	}
	json.NewEncoder(w).Encode(users)
}

// Создание нового пользователя через SQL
func createUserSQL(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err = dbSQL.Exec("INSERT INTO users (name, age) VALUES ($1, $2)", user.Name, user.Age)
	if err != nil {
		http.Error(w, "Could not insert user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func updateUserSQL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	_, err = dbSQL.Exec("UPDATE users SET name = $1, age = $2 WHERE id = $3", user.Name, user.Age, userID)
	if err != nil {
		http.Error(w, "Could not update user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

func deleteUserSQL(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	_, err := dbSQL.Exec("DELETE FROM users WHERE id = $1", userID)
	if err != nil {
		http.Error(w, "Could not delete user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
}

// Получение всех пользователей через GORM с фильтрацией и сортировкой
func getUsersGORM(w http.ResponseWriter, r *http.Request) {
	var users []User
	query := dbGORM

	ageFilter := r.URL.Query().Get("age")
	if ageFilter != "" {
		age, _ := strconv.Atoi(ageFilter)
		query = query.Where("age >= ?", age)
	}

	sortParam := r.URL.Query().Get("sort")
	if sortParam == "name" {
		query = query.Order("name ASC")
	}

	query.Find(&users)
	json.NewEncoder(w).Encode(users)
}

func createUserGORM(w http.ResponseWriter, r *http.Request) {
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	result := dbGORM.Create(&user)
	if result.Error != nil {
		http.Error(w, "Could not insert user", http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusCreated)
}

func updateUserGORM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil || user.Name == "" {
		http.Error(w, "Invalid input", http.StatusBadRequest)
		return
	}

	dbGORM.Model(&User{}).Where("id = ?", userID).Updates(User{Name: user.Name, Age: user.Age})
	w.WriteHeader(http.StatusOK)
}

func deleteUserGORM(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userID, _ := strconv.Atoi(vars["id"])

	dbGORM.Delete(&User{}, userID)
	w.WriteHeader(http.StatusOK)
}

// Основная функция
func main() {
	// Подключение к базе данных
	dbSQL = connectSQL()
	dbGORM = connectGORM()

	// Настройка маршрутов
	r := mux.NewRouter()

	r.HandleFunc("/sql/users", getUsersSQL).Methods("GET")
	r.HandleFunc("/sql/users", createUserSQL).Methods("POST")
	r.HandleFunc("/sql/users/{id}", updateUserSQL).Methods("PUT")
	r.HandleFunc("/sql/users/{id}", deleteUserSQL).Methods("DELETE")

	r.HandleFunc("/gorm/users", getUsersGORM).Methods("GET")
	r.HandleFunc("/gorm/users", createUserGORM).Methods("POST")
	r.HandleFunc("/gorm/users/{id}", updateUserGORM).Methods("PUT")
	r.HandleFunc("/gorm/users/{id}", deleteUserGORM).Methods("DELETE")

	fmt.Println("Server is running on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}
