package main

import (
	"fmt"
	"time"
)

func main() {
	language := "Golang"

	switch language {
	case "Python":
		fmt.Println("Good, go for Go! You are using")
	case "Go", "golang":
		fmt.Println("Good, go for Go! You are using curly braces")
	default:
		fmt.Println(" Any other programming language is a good start!")
	}

	n := 5

	switch {
	case n%2 == 0:
		fmt.Println("Even number!")
	case n%2 != 0:
		fmt.Println("Odd number!")
	default:
		fmt.Println("Never here!")
	}

	hour := time.Now().Hour()
	fmt.Println(hour)
	switch true {
	case hour < 12:
		fmt.Println("Good morning")
	case hour < 17:
		fmt.Println("Gopod afternoon")
	default:
		fmt.Println("Good evening")
	}
}
