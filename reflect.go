package main
import (
	"fmt"
	"reflect"
)

type Person struct {
	Name string `: required`
	Age  int	`: # >= 0`
}

func main() {
	var haren = Person{Name: "Haren Lin", Age: 22}
	var harenPtr = &haren

	typeOfHaren := reflect.TypeOf(haren)
	typeOfHarenPtr := reflect.TypeOf(harenPtr)
	
	fmt.Println("Name:", typeOfHaren.Name(), "Kind:", typeOfHaren.Kind()) // Name: Person Kind: struct
	fmt.Println("Name:", typeOfHarenPtr.Name(), "Kind:", typeOfHarenPtr.Kind()) // Name:  Kind: ptr

	for i := 0; i < typeOfHaren.NumField(); i++ {
		fmt.Printf("%+v\n", typeOfHaren.Field(i))
	}
	// {Name:Name PkgPath: Type:string Tag: Offset:0 Index:[0] Anonymous:false}
	// {Name:Age PkgPath: Type:int Tag: Offset:16 Index:[1] Anonymous:false}

	for i := 0; i < typeOfHaren.NumField(); i++ {
		fmt.Printf("%v %v\n", typeOfHaren.Field(i).Name, typeOfHaren.Field(i).Tag)
	}
	// Name : required
	// Age : # >= 0
}