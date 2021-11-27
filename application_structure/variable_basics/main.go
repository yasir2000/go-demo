package main

import "fmt"

func main() {
	var age = 30
	fmt.Println("Age:", age)

	var name = "Dan"
	fmt.Println("Your name is :", name)
	_ = name

	s := "learning golang"
	fmt.Println(s)

	name = "Andrei"
	name1 := "John"
	_ = name1

	car, cost := "Audi", 50000
	fmt.Println(car, cost)
	car, year := "BMW", 2018
	_ = year

	var opened = false
	opened, file := true, "a.txt"
	_, _ = opened, file

	var (
		salary    float64
		firstName string
		gender    bool
	)
	fmt.Println(salary, firstName, gender)

	var a, b, c int
	fmt.Println(a, b, c)

	var i, j int
	i, j = 5, 8

	j, i = i, j //swapping variables
	fmt.Println(i, j)

	sum := 5 + 2.3
	fmt.Println(sum)

	var d = 4
	var e = 5.2
	d = int(e)
	fmt.Println(d, e)

	var value int
	var prince float64
	var name2 string
	var done bool

	fmt.Println(value, prince, name2, done)

	name3 := "Andrie"
	fmt.Println("Hellow, playground. I am", name3)
	a1, b1 := 4, 6
	fmt.Println("Sum:", a1+b1, "Mean Value:", (a1+b1)/2)

	fmt.Printf("Your age is %d\n", 21)
	fmt.Printf("x is %d, y is %f", 5, 6.8)
	fmt.Printf("He says : \"Hellow Go!\"\n")

	figure := "Circle"
	radius := 5
	pi := 3.14159

	_, _ = figure, pi
	fmt.Printf("Radius is %d\n", radius)
	fmt.Printf("Radius is %+d\n", radius)
	fmt.Printf("Pi constant is %f\n", pi)
	fmt.Printf("The area of a %s with a Radius of %d is %f\n", figure, radius, float64(radius/2)*2*pi)
	fmt.Printf("The perimiter of a %s with Radius of %d is %f\n", figure, radius, float64(2*radius)*2*pi)

	//%q for quoted string
	fmt.Printf("This is a %q\n", figure)
	fmt.Printf("This is a %s", figure)

	//%v -> replaced by any type
	fmt.Printf("The area of a %v with a Radius of %v is %v\n", figure, radius, float64(radius/2)*2*pi)

	//%T -> type of a variable
	fmt.Printf("The area of a %v %T with a Radius of %v %T is %v %T\n", figure, radius, float64(radius/2)*2*pi, figure, radius, float64(radius/2)*2*pi)

	//%t -> bool
	closed := true
	fmt.Printf("Flie closed: %t\n", closed)

	//%b -> base 2
	fmt.Printf("%08b \n", 55)

	//%x -> hex 16

	x := 3.4
	y := 6.9
	fmt.Printf("x * y = %.3f\n", x*y)

}
