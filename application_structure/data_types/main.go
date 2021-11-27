package main

import "fmt"

func main() {

	// INT TYPE
	var i1 int8 = -128
	fmt.Printf("%T\n", i1)

	var i2 uint16 = 65535
	fmt.Printf("%T\n", i2)

	var f1, f2, f3 float64 = 1.1, -.2, 5.
	fmt.Printf("%T %T %T\n", f1, f2, f3)

	var r rune = 'f'
	fmt.Printf("%T\n", r)
	fmt.Println(r)
	fmt.Printf("%x\n", r)

	var b bool = true
	fmt.Printf("%T\n", b)

	var s string = "Hello Go!"
	fmt.Printf("%T\n", s)

	var numbers = [4]int{4, 5, -9, 100}
	fmt.Printf("%T\n", numbers)

	var cities = []string{"London", "Tokyo", "New york"}
	fmt.Printf("%T\n", cities)

	balances := map[string]float64{
		"USD": 2322.2,
		"EUR": 511.22,
	}
	fmt.Printf("%v %T\n", balances, balances)

	type Person struct {
		name string
		age  int
	}

	var you Person
	fmt.Printf("%T\n", you)

	var x int = 22
	ptr := &x
	fmt.Printf("ptr is of type %T with the value of %v\n", ptr, ptr)

	fmt.Printf("%T\n", f)

}
func f() {

}
