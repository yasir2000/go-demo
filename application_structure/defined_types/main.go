package main

import (
	"fmt"
	"os"
)

type km float64
type mile float64

func main() {
	type speed uint
	var s1 speed = 10
	var s2 speed = 20

	fmt.Println(s2 - s1)

	var x uint
	x = uint(s1)
	_ = x

	var s3 speed = speed(x)

	_ = s3

	var parisToLondon km = 465
	var distanceInMiles mile

	distanceInMiles = mile(parisToLondon) / 0.621
	fmt.Println(distanceInMiles)

	type second = uint

	var hour second = 3600
	fmt.Printf("Minutes in an hour : %d \n", hour/60)

	price, inStock := 100, true

	if price > 80 {

		fmt.Println("Too Expensive!")
	}

	_ = inStock

	if price <= 100 && inStock == true {
		fmt.Println("Buy it!")
	}

	lowPrice := true
	if lowPrice {
		fmt.Println(("something"))
	}

	fmt.Println("os.Args", os.Args)

}
