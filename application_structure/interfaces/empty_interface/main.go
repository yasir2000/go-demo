package main

import "fmt"

//type emptyInterface interface {
//}

type person struct {
	info interface{}
}

func main() {
	var empty interface{}
	empty = 5
	fmt.Println(empty)

	empty = "Go"
	fmt.Println(empty) // assertion

	empty = []int{4, 5, 6}
	//fmt.Println(len(empty))
	// we cannot check len of empty interface
	// we cannot use directly empty interface in operation
	// an interface stores two values; a dynamic type and a dynamic concerete value
	// to access the dynamic value, we have to do a type assertion
	// fro above example, if we want to use a function that accepts a slice as argument
	// or a method that works on slice value, we must retrieve the dynamic value using an assertion
	fmt.Println(empty.([]int))      //assertion but we had to put (.) because its a slice assertion
	fmt.Println(len(empty.([]int))) //len of interface after getting its dynamic concrete value using assertion
	// empty interfaces are useful to handle values of unkown type
	// you can pass an empty interface type as a function parameter of any type
	you := person{}
	you.info = "Your name"
	you.info = 40
	you.info = []float64{5.6, 5., 7.8}
	fmt.Println(you.info)

	// one of the defects of using empty interfaces, is that type checking during runtime
	// can be difficult, to trace unkown types, so we need to use empty interfaces
	// only if its neccessary
}
