package main
import "fmt"

type Person struct{
	Name string
	Age int
}

// func changePerson(p *Person, nName string, nAge int) {
// 	p.Name = nName
// 	p.Age = nAge
// }

func main() {
	// var intern = &Person{"Kowala", 18}
	// fmt.Println("Before change:", *intern) // Before change: {Kowala 18}
	// changePerson(intern, "Haren Lin", 22)
	// fmt.Println("After change:", *intern)  // After change: {Haren Lin 22}
	var intern = new(Person)
	intern.Name = "Haren Lin"
	intern.Age = 22
	fmt.Println(*intern) // {Haren Lin 22}
}