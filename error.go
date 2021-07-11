package main
import "fmt"
import "errors"

func divide(dividend, divisor int) (int, error) { 
	if divisor == 0 { // error occurs when divisor is 0
		return 0, errors.New("Hey! A number can't divided by zero!")
	}
	return dividend / divisor, nil // error is nil while running successfully
}

func main() {
	value, error := divide(100, 0)
	if error == nil {
		fmt.Println(value)
	} else {
		fmt.Println(error)
	}
}