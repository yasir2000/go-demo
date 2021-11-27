package main

import "fmt"

func main() {
	// 	const a float64 = 5.1
	// 	const b = 6.7
	// 	const c float64 = a * b
	// 	const str = "Hello" + "Go!"
	// 	const d = 5 > 10
	// 	fmt.Println(d)

	// 	//const x int = 5
	// 	//	const y float64 = 2.2 * x

	const (
		c1 = (iota * 2) + 1
		c2 = (iota + 3) * 2
		c3 = (iota + c2) * 3
	)
	fmt.Println(c1, c2, c3)
	const (
		c11 = iota + c1
		c12 = iota + c2
		c13 = iota + c11 + c2
	)

	fmt.Println(c11, c12, c13)

}
