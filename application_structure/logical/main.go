package main

import (
	"fmt"
)

func main() {

	a, b := 5, 10
	fmt.Println(a == b)
	fmt.Println(a != b)
	fmt.Println(a > b, a >= b)
	fmt.Println(b < a || 10 <= b)

	fmt.Println(a > 1 && b == 10)

}
