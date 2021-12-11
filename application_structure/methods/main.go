package main

import (
	"fmt"
	"time"
)

type names []string

func (n names) print() {
	for i, name := range n {
		fmt.Println(i, name)
	}
}

type car struct {
	brand string
	price int
}

func changeCar(c car, newBrand string, newPrice int) {
	c.price = newPrice
	c.brand = newBrand
}

func (c car) changeCar1(newBrand string, newPrice int) {
	c.brand = newBrand
	c.price = newPrice
}

func (c *car) changeCar2(newBrand string, newPrice int) {
	(*c).brand = newBrand
	(*c).price = newPrice
}

// type distance *int

// func (d *distance) m1 (){ // not allowed to create a method on pointer types, like *int
// 	fmt.Println("Just a message")
// }

func main() {
	const day = 24 * time.Hour
	fmt.Printf("%T\n", day)

	seconds := day.Seconds()
	fmt.Printf("%T\n", seconds)

	fmt.Printf("Seconds in a day: %v\n", seconds)

	friends := names{"Dan", "Marry"}
	friends.print()

	names.print(friends)

	var n int64 = 9273636346
	fmt.Println(n)
	fmt.Println(time.Duration(n))

	myCar := car{brand: "Audi", price: 40900}
	changeCar(myCar, "Renault", 20000)
	fmt.Println(myCar)

	myCar.changeCar1("Opel", 20001222)
	fmt.Println(myCar)

	(&myCar).changeCar2("Seat", 87777)
	fmt.Println(myCar)

	var yourCar *car
	yourCar = &myCar
	fmt.Println(*yourCar)

	//valid ways
	yourCar.changeCar2("VW", 30023)
	fmt.Println(*yourCar)

	(*yourCar).changeCar2("Another VW", 3230)
	fmt.Println(*yourCar)

	fmt.Println(myCar)
}
