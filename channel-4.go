package main
import "fmt"

func square(s []int, chnSendOnly chan<- int) {
	for _, val := range s {
		chnSendOnly <- val * val
	}
	close(chnSendOnly) // close channel & tell range for loop is done
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	chn := make(chan int)
	var chnSendOnly chan<- int = chn // create single directional channel
	go square(arr, chnSendOnly)
	for squared_val := range chn {
		fmt.Println(squared_val)
	}
}