package main
import "fmt"

type Person struct {
	Name string
}

func (p *Person) says (said string) {
	fmt.Println(said)
}

type Member struct {
	// Person   *Person // struct in struct
	*Person  // 不需要欄位名稱
	TeamName string
	TeamID   int
}

func (m *Member) code (language string) {
	fmt.Printf("I can code with programming language: %s\n", language)
}

func main() {
	var haren = &Member{&Person{"Haren Lin"}, "ERS", 16}
	haren.says("Kowala") // = haren.Person.says("Kowala")
	haren.code("C++, Python, Go") // I can code with programming language: C++, Python, Go
}