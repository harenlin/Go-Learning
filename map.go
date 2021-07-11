package main
import "fmt"
func main() {
    var operation = make(map[string]func(a, b int) int)
    operation["add"] = func(a, b int) int {
        return a + b
    }
    operation["minus"] = func(a, b int) int {
        return a - b
    }
    operation["multiply"] = func(a, b int) int {
        return a * b
    }
    operation["divide"] = func(a, b int) int {
        return a / b
    }
    
    for k, v := range operation {
        fmt.Printf("%s(3,9) = %d \n", k, v(3, 9))
    }
}

/* print out:
divide(3,9) = 0 
add(3,9) = 12 
minus(3,9) = -6 
multiply(3,9) = 27 
*/