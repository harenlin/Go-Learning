package main
import (
	"fmt"
	"time"
)

func sayHello(id string) {
	for idx := 1; idx < 5; idx++ {
		fmt.Println("hello from: ", id, " - ", idx)
		time.Sleep(time.Second)
	}
}

func main() {
	go sayHello("A")
	go sayHello("B")
	go sayHello("C")
	for idx := 1; idx < 5; idx++ {
		fmt.Println("hello from:  main  - ", idx)
		time.Sleep(time.Second)
	}
	time.Sleep(3 * time.Second)
}