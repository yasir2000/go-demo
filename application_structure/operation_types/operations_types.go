package main

import (
	"fmt"
)

func main() {
	a, b := 4, 20
	r := (a + b) / (a - b) * 2

	fmt.Println(r)

	r = 9 % a

	b -= 2
	fmt.Println(b)

	b *= 10
	fmt.Println(b)

	b /= 5
	fmt.Println(b)

	b %= 3
	fmt.Println(b)
}
