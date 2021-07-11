package main
import "fmt"

func printSlice(slice []string) {
	fmt.Printf("len = %d, capacity = %d, values= ", len(slice), cap(slice))
	for k, v := range slice {
		fmt.Printf("[%d]%s, ", k, v)
	}
	fmt.Println()
}

func main() {
	slice := make([]string, 0)
	slice = append(slice, "NTU") // dynamic allocate append a string
	printSlice(slice)
	slice = append(slice, "NCCU") // dynamic allocate append a string
	printSlice(slice)
	slice = append(slice, "NCKU") // dynamic allocate append a string
	printSlice(slice)
}

/* print out:
len = 1, capacity = 1, values= [0]NTU, 
len = 2, capacity = 2, values= [0]NTU, [1]NCCU, 
len = 3, capacity = 4, values= [0]NTU, [1]NCCU, [2]NCKU, 
*/