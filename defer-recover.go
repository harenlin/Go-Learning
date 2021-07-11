package main
import "fmt"

func main() {
	fmt.Println("Taiwan: Here Comes 2021")

	defer func() {
		if error := recover(); error != nil {
			fmt.Println(error)
			fmt.Println("Don't panic! We got vaccinated!")
		}
		fmt.Println("Life still moves on!")
	}()

	panic("Oh! No! COVID Pandemic!")
}