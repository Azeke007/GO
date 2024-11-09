package main

import "fmt"

type Employee struct {
    Name string
    ID   int
}

type Manager struct {
    Employee
    Department string
}

func (e Employee) Work() {
    fmt.Printf("Employee %s with ID %d is working.\n", e.Name, e.ID)
}

func main() {
    manager := Manager{
        Employee: Employee{Name: "John Doe", ID: 123},
        Department: "Sales",
    }
    manager.Work()
}