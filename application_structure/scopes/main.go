package main

import (
	"fmt"
	f "fmt"
)

const done = false

var b int = 10

func main() {
	var task = "Running" //
	f.Println(task, done)

	const done = true
	f.Printf("done in main() is v%\n", done)
	f1()
	f.Println("Bye Bye!")
}

func f1() {
	const done = true
	f.Printf("done in f1(): %v\n", done)

	a := 10
	_ = a
	//count := 0
	for i := 1; i <= 500; i++ {
		if i%7 == 0 && i%5 == 0 { // if i is divisible by 7 and 5
			fmt.Printf("%d ", i)
			//break

		}

		//count++
		//if count == 3 {
		//	break
		//}
	}
	fmt.Println("")
	birthYear, currentYear := 1980, 2020

	for i := birthYear; i <= currentYear; {
		fmt.Printf("%d ", i)
		i++
	}
	fmt.Println()

	age := 10
	switch {
	case age < 0 || age > 100:
		fmt.Println("Invalid Age")
	case age < 18:
		fmt.Println("You are minor!")
	case age == 18:
		fmt.Println("Congratulations! You've just become major!")
	default:
		fmt.Println("You are major!")
	}
}
