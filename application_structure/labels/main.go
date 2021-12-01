package main

import (
	"fmt"
)

func main() {

	outer := 19
	_ = outer
	people := [5]string{"Hellen", "Mark", "Brenda", "Antonio", "Micheal"}
	friends := [2]string{"Mark", "Merry"}

outer:
	for index, name := range people {
		for _, friend := range friends {
			if name == friend {
				fmt.Printf("Found a friend %q at index %d\n", friend, index)
				break outer
			}
		}
	}
	fmt.Println("Next instruction after the break")

	i := 0

loop:
	if i < 5 {
		fmt.Println(i)
		i++
		goto loop

	}

}
