package main

import (
	"fmt"
	"strings"
)

func main() {
	var numbers [4]int

	fmt.Printf("%v\n", numbers)
	fmt.Printf("%#v\n", numbers)

	var a1 = [4]float64{}
	fmt.Printf("%#v\n", a1)

	var a2 = [3]int{-10, 1, 100}
	fmt.Printf("%#v\n", a2)

	a3 := [4]string{"dan", "diana", "paul", "john"}
	fmt.Printf("%#v\n", a3)

	a4 := [4]string{"x", "y"}
	fmt.Printf("%#v\n", a4)

	a5 := [...]int{1, 2, 3, 45, 56}
	fmt.Printf("%#v\n", a5)
	fmt.Printf("the length of a5 is %d\n", len(a5))

	a6 := [...]int{1, 2, 5, 182, 77}
	fmt.Printf("%#v\n", a6)

	numbers[0] = 77
	fmt.Printf("%#v\n", numbers)

	numbers[3] = 100
	fmt.Printf("%#v\n", numbers)

	for i, v := range numbers {
		fmt.Println("index:", i, " value:", v)
	}

	fmt.Println(strings.Repeat(" Yasir ", 20))
	for i := 0; i < len(numbers); i++ {
		fmt.Printf("index:\n", i, "value:", numbers[i])
	}
	fmt.Printf("\n")
	balances := [3][3]int{
		{5, 6, 7},
		{8, 9, 10},
		{8, 1, 3},
	}
	fmt.Println(balances)

	m := [3]int{1, 2, 3}
	n := m
	fmt.Println("n is equal to m:", n == m)
	m[0] = -10
	fmt.Println("n is equal to m:", n == m)

	grades := [3]int{
		1: 10,
		0: 10,
		2: 7,
	}
	fmt.Println(grades)

	accounts := [3]int{2: 50}
	fmt.Println(accounts)

	names := [...]string{
		5: "Dan",
	}
	fmt.Println(names, len(names))

	cities := [...]string{
		5: "Paris",
		"London",
		1: "NYC",
	}
	fmt.Printf("%#v\n", cities)

	weekend := [7]bool{5: true, 6: true}
	fmt.Println(weekend)

	var cities1 [2]string
	fmt.Printf("%#v\n", cities1)

	// 2.
	grades1 := [3]float64{4.5, 9.7}
	fmt.Printf("%#v\n", grades1)

	// 3.
	taskDone := [...]bool{true, false, false, true}
	fmt.Printf("%#v\n", taskDone)

	// 4.
	for i := 0; i < len(cities1); i++ {
		fmt.Println(cities1[i])
	}

	// 5.
	for index, value := range grades1 {
		fmt.Println(index, value)
	}

	nums := [...]int{30, -1, -6, 90, -6}
	for _, value := range nums {
		if value >= 0 && value%2 == 0 {
			fmt.Println(value)
		}
	}
	myArray := [4]float64{1.2, 5.6}

	myArray[2] = 6

	a := 10
	myArray[0] = float64(a)

	myArray[3] = 10.10

	fmt.Println(myArray)
}
