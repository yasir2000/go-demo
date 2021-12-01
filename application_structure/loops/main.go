package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	j := 10

	for j >= 0 {
		fmt.Println(j)
		j--
	}

	for i := 0; i < 10; i++ {
		if i == 5 {
			break
		}
		fmt.Println(i)
	}

	count := 0

	for x := 0; true; x++ {
		if x%13 == 0 {
			fmt.Printf("%d is divisable by 13 \n", x)
			count++
		}
		if count == 10 {
			break
		}
	}
	fmt.Println("Just a message")

}
