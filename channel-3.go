package main
import "fmt"

func main() {
	chn := make(chan int, 2) // declare a channel with buffer len = 2
    fmt.Println(len(chn))    // use len(chn) get the # of data in the buffer 
	chn <- 45
	fmt.Println(len(chn))
	chn <- 50
	fmt.Println(len(chn))
}