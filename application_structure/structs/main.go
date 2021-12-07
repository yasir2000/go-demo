package main

import "fmt"

func main() {
	title1, author1, year1 := "The Devine Comedy", "Dante Aligheri", 1320
	title2, author2, year2 := "Macbeth", "William Shakespeare", 1606

	fmt.Println("Book1", title1, author1, year1)
	fmt.Println("Book1", title2, author2, year2)

	type book struct {
		title  string
		author string
		year   int
	}

	type book1 struct {
		title, author string
		year          int
	}

	myBook := book{"The Dvine Comdey", "Dante Aligheri", 1320}
	fmt.Println(myBook)

	bestBook := book{title: "Animal Farm", year: 1945, author: "George Orwell"}
	fmt.Println(bestBook)

	aBook := book{title: "Just a random book!"}
	fmt.Printf("%#v\n", aBook)

	lastBook := book{title: "Anna Karenina"}
	fmt.Println(lastBook.title)

	fmt.Printf("%#v\n", lastBook)
	lastBook.author = "Leo Tolestoy"
	fmt.Printf("%+v\n", lastBook)

	aBook1 := book{title: "Anna Karenina", author: "Micheal", year: 1878}
	fmt.Println(aBook1 == lastBook)

	myBook1 := aBook
	myBook1.year = 2020
	fmt.Println(myBook1, aBook)

	diana := struct {
		firstName, lastName string
		age                 int
	}{
		firstName: "Diana",
		lastName:  "Muller",
		age:       30,
	}
	fmt.Printf("%#v\n", diana)
	fmt.Printf("Diana's Age: %d\n", diana.age)

	type Book2 struct {
		string
		float64
		bool
	}

	b1 := Book2{"1984 by George Orwell", 10.2, false}

	fmt.Printf("%#v\n", b1)

	fmt.Println(b1.string)

	type Employee struct {
		name   string
		salary int
		bool
	}

	e := Employee{"John", 4000, false}
	fmt.Printf("%#v\n", e)
	e.bool = true
	fmt.Printf("%#v\n", e.bool)

}
