package main

import (
	"fmt"
)

func main() {
	type book struct {
		id     string
		title  string
		author string
	}

	type library struct {
		id       string
		location [2]float64
		books    map[string]book
	}

	jonstonPublic := library{
		id:       "aaaa",
		location: [2]float64{20.32, 22.33},
		books:    map[string]book{},
	}

	fmt.Println("Empty Library:", jonstonPublic)

	b1 := book{
		id:     "bbbb",
		title:  "Where the red fern grows",
		author: "Wilson Rawls",
	}
	b2 := book{
		id:     "cccc",
		title:  "All is quiet on the western front",
		author: "Erich Maria Remarque",
	}
	b3 := book{
		id:     "dddd",
		title:  "Catch 22",
		author: "Joseph Heller",
	}

	jonstonPublic.books[b1.id] = b1
	jonstonPublic.books[b2.id] = b2
	jonstonPublic.books[b3.id] = b3

	fmt.Printf("Library with 3 books: %#v\n\n", jonstonPublic)

	// Trying to change the title of b3 ("Catch 22" to "The Book Thief")
	// Gives compiler error: [cannot assign to struct field jonstonPublic.books[b3.id].title in map]
	//jonstonPublic.books[b3.id].title = "The Book Thief"
	// First we get a "copy" of the entry
	if entry, ok := jonstonPublic.books[b3.id]; ok {

		// Then we modify the copy
		entry.id = "The Book Thief"

		// Then we reassign map entry
		jonstonPublic.books[b3.id] = entry

	}

	// Now b3.id key. books is "The Book Thief"
	fmt.Println(jonstonPublic.books)

}
