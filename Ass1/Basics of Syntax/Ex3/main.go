// Control Structures
package main

import (
    "fmt"
)

func main(){
    var num int
    fmt.Print("Enter an integer: ")
    fmt.Scan(&num)

    if num > 0 {
        fmt.Println("The number is positive.")
    } else if num < 0 {
        fmt.Println("The number is negative.")
    } else {
        fmt.Println("The number is zero.")
    }

	sum := 0
    for i := 1; i <= 10; i++ {
        sum += i
    }
    fmt.Printf("The sum of the first 10 natural numbers is: %d\n", sum)

    var day int
    fmt.Print("Enter a number (1 for Monday, 2 for Tuesday, ... 7 for Sunday): ")
    fmt.Scan(&day)

    switch day {
    case 1:
        fmt.Println("Monday")
    case 2:
        fmt.Println("Tuesday")
    case 3:
        fmt.Println("Wednesday")
    case 4:
        fmt.Println("Thursday")
    case 5:
        fmt.Println("Friday")
    case 6:
        fmt.Println("Saturday")
    case 7:
        fmt.Println("Sunday")
    default:
        fmt.Println("Invalid input! Please enter a number between 1 and 7.")
    }
}
