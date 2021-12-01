package main

import "fmt"

func main() {
	var cities []string
	fmt.Println("cities is equal to nil : ", cities == nil)
	fmt.Printf("cities %#v\n", cities)
	fmt.Println(len(cities))

	cities[0] = "London"

}
