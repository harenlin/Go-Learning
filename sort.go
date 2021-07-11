package main
import "fmt"
import "sort"

type Person struct {
    Name string
    Age  int
}

// type Arr_P []Person
// sort.Sort is a interface which has three methods: Len, Swap, Less
// we have to define it in detailed by ourselves
// func (a Arr_P) Len() int           { return len(a) }
// func (a Arr_P) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
// func (a Arr_P) Less(i, j int) bool { return a[i].Age < a[j].Age }

func (a []Person) Len() int           { return len(a) }
func (a []Person) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a []Person) Less(i, j int) bool { return a[i].Age < a[j].Age }

func main() {
    people := []Person{{"Bob", 31}, {"John", 42}, {"Michael", 17}, {"Jenny", 26}}
    fmt.Println(people) // [{Bob 31} {John 42} {Michael 17} {Jenny 26}]
    sort.Sort(Arr_P(people))
    fmt.Println(people) // [{Michael 17} {Jenny 26} {Bob 31} {John 42}]
}
