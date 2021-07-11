package main
import "fmt"

// func player(p_name string, p_chan chan int) {

// }

// func main() {
// 	fmt.Println("Hello Go!")
// }

func sum(arr []int, chn chan int) {
	sum := 0
	for _, v := range arr {
		sum += v
	}
	chn <- sum // 傳送總和
}

func main() {
	arr := []int{1, 2, 3, 4, 5, 6}
	chn1 := make(chan int)
	chn2 := make(chan int)
	go sum(arr[:4], chn1)
	go sum(arr[4:], chn2)
	x, y := <-chn1, <-chn2 // receive sum
	fmt.Println(x, y, x+y) // 10 11 21
}
