package main
import "fmt"

func swap(a string, b string) (string, string) {
    return b, a
}

func main() {
    str1 := "Hello"
    str2 := "World"
    fmt.Println("Before swap:", str1, str2)
    str1, str2 = swap(str1, str2)
    fmt.Println("After swap:", str1, str2)
}