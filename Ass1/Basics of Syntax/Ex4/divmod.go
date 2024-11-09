package main
import "fmt"

func divmod(a int, b int) (int, int) {
    return a / b, a % b
}

func main() {
    quotient, remainder := divmod(17, 5)
    fmt.Println("Quotient:", quotient)
    fmt.Println("Remainder:", remainder)
}