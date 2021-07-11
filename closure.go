package main
import "fmt"

type counter struct {
	add   func()
	minus func()
	print func()
}

func main() {
	count := func() counter {
		idx := 0
		return counter{
			func() { idx++ },
			func() { idx-- },
			func() { fmt.Println("idx =", idx) },
		}
	}()
	
	count.add()
	count.print() // idx = 1
	count.add()
	count.print() // idx = 2
	count.minus()
	count.print() // idx = 1
}