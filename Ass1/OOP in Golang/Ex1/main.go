package main
import "fmt"

type Person struct {
    Name string
    Age  int
}

func (p Person) Greet() {
    fmt.Printf("Hello, my name is %s and I am %d years old.\n", p.Name, p.Age)
}

func main() {
    person := Person{Name: "John", Age: 30}
    person.Greet()
}