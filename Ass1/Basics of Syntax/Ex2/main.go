// Variables and Data Types
package main

import "fmt"

func main() {
    var a int = 10
    var b float64 = 20.5
    var c string = "Hello, World!"
    var d bool = true

    fmt.Printf("The value of a is: %d and its type is: %T\n", a, a)
    fmt.Printf("The value of b is: %f and its type is: %T\n", b, b)
    fmt.Printf("The value of c is: %s and its type is: %T\n", c, c)
    fmt.Printf("The value of d is: %t and its type is: %T\n", d, d)
}