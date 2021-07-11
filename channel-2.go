package main
import "fmt"

func square(s []int, chn chan int) {
	for _, val := range s {
		chn <- val * val
	}
	close(chn) // close channel & tell range for loop is done
}

func main() {
	arr := []int{1, 3, 5, 7, 9}
	chn := make(chan int)
	go square(arr, chn)
	for squared_val := range chn {
		fmt.Println(squared_val)
	}
}