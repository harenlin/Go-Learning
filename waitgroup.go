package main
import (
	"fmt"
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
	var wait_group sync.WaitGroup // declare WaitGroup variable to wait all tasks done
	for i := 0; i < 500; i++ {
		wait_group.Add(1) // add a new task
		go func() {
			defer wait_group.Done() // wait_group.Done()成立時，代表這個 func() 已經走完，也代表 addCount 完成
			addCount()
		}()
	}
	wait_group.Wait() // waiting all of the tasks done
	fmt.Println("cnt = ", getGount())
}