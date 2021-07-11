package main
import (
	"fmt"
	"time"
)

var cnt int
func addCount() {
	cnt++
}
func getGount() int {
	return cnt
}

func main() {
	for i := 0; i < 500; i++ {
		go addCount()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("cnt = ", getGount())
}