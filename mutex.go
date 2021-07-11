package main
import (
	"fmt"
	"time"
	"sync"
)

var cnt int
var mutex_lock sync.Mutex

func addCount() {
	mutex_lock.Lock()
	defer mutex_lock.Unlock()
	cnt++
}
func getGount() int {
	mutex_lock.Lock()
	defer mutex_lock.Unlock()
	return cnt
}

func main() {
	for i := 0; i < 500; i++ {
		go addCount()
	}
	time.Sleep(1 * time.Second)
	fmt.Println("cnt = ", getGount())
}