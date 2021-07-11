package main
import "fmt"
// func swap(a, b *int) {
// 	// temp := *a
//     var temp int = *a
// 	*a = *b
// 	*b = temp
// }

func main() {
	// var ptr *string  // 宣告一個指向字串的指標變數ptr
	// fmt.Printf("%p \n", ptr)  // 印出: 0x0 (尚未指派記憶體位置)
	// if ptr == nil {
	//     fmt.Println("ptr is nil")
	// }
	// str := "Haren Lin"
	// var ptr *string = &str
	// fmt.Printf("%p \n", ptr)  // print out: 0xc000096220
	// fmt.Printf("%s \n", *ptr) // print out: Haren Lin
    // a, b := 54, 39
	// swap(&a, &b)
	// println(a, b)
    ptr := new(string)
    *ptr = "Haren Lin"
    fmt.Printf("%s \n", *ptr)
}
