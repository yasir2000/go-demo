package main

import "fmt"

func main() {
	var employees map[string]string
	fmt.Printf("%#v\n", employees)

	fmt.Printf("No of pairs : %d\n", len(employees))
	fmt.Printf("The value for key %q is %q", "Dan", employees["Dan"])

	var accounts map[string]float64
	fmt.Printf("%#v\n", accounts["0x323"])

	var m1 map[[3]int]string
	_ = m1

	//employees["Dan"] = "Programmer"

	people := map[string]float64{}

	people["John"] = 21.4
	people["Marry"] = 12.3
	fmt.Println(people)

	map1 := make(map[int]int)
	map1[4] = 6

	balances := map[string]float64{
		"USD": 332.11,
		"EUR": 233.12,
		"TRY": 12.12,
	}
	fmt.Println(balances)

	m := map[int]int{1: 10, 2: 21, 3: 21}
	_ = m

	balances["USD"] = 43.45
	balances["GBP"] = 288.32
	fmt.Println(balances)

	fmt.Println(balances["RON"])

	balances["RON"] = 0

	v, ok := balances["RON"]
	if ok {
		fmt.Println("the RON balance is :", v)

	} else {
		fmt.Println("The RON Key doesnt exist in the map!")
	}

	// declaring a map with keys of type string and values of type string
	var employees1 map[string]string
	//the zero value of a map is nil

	// the zero value of a map is nil
	fmt.Printf("%#v\n", employees1) // -> map[string]string(nil).

	fmt.Printf("No. of elements: %d\n", len(employees1)) // => No. of elements: 0

	// getting the corresponding value of a key
	// if the key doesn't exist or the map is not initialized it returns the zero value for the value type
	fmt.Printf("The value for key %q is %q \n", "Dan", employees1["Dan"]) // => The value for key "Dan" is ""

	// key must be of comparable types
	// var m1 map[[]int]string // error -> invalid map key type []int (slices are not comparable)

	// inserting a key:value pair in a nil map
	// employees["Dan"] = "Programmer" // error -> panic: assignment to entry in nil map

	// declaring and initializing a map using a map literal
	people1 := map[string]float64{} // empty map

	// inserting key:value pairs in a map
	people1["John"] = 30.5
	people1["Marry"] = 22

	fmt.Printf("%v\n", people1) // => map[John:30.5 Marry:22]

	// declaring and initializing a map using the make() function:
	map11 := make(map[int]int)
	fmt.Printf("map1: %#v\n", map11) // -> map1: map[int]int{}
	map1[4] = 8

	// declaring and initializing a map using a map literal
	balances1 := map[string]float64{
		"USD": 233.11,
		"EUR": 555.11,
		//50: "ABC"  //illegal, all keys have the same type which is string and all values have the same type which is float64
		"CHF": 600, //this last comma (,) is mandatory when declaring the map on multiple lines
	}
	fmt.Println(balances1) // => map[CHF:600 EUR:555.11 USD:233.11]

	//If we declare and initialize a map on a single line, it's not mandatory to use a comma after the last element
	m11 := map[int]int{1: 3, 4: 5, 6: 8}
	_ = m11

	// initializing a map with duplicate keys
	// n := map[int]int{1: 3, 4: 5, 6: 8, 1: 4} // error -> duplicate key 1 in map literal

	// if the key exists it updates its value and if the key doesn't exist it adds the key: value pair
	balances1["USD"] = 500.5
	balances1["GBP"] = 800.8
	fmt.Println(balances1) // => map[CHF:600 EUR:555.11 GBP:800.8 USD:500.5]

	// "comma ok" idiom is used to distinguish between a missing key:value pair and an existing key with value zero
	v1, ok := balances1["RON"]

	//v is the key's corresponding value
	// ok is bool type value which is true if the key exists and false otherwise

	if ok {
		fmt.Println("The RON Balance is: ", v1)
	} else {
		fmt.Println("The RON key doesn't exist in the map!")
	}

	// iterating over a map
	for k, v1 := range balances1 {
		fmt.Printf("Key: %#v, Value: %#v\n", k, v1)
	}

	//starting with go 1.12 fmt.Printf() function prints out the map sorted by key.
	fmt.Printf("balances: %v\n", balances1) // => balances: map[CHF:600 EUR:555.11 GBP:800.8 USD:500.5]

	// deleting a key:value pair from the map
	delete(balances1, "USD")

	//** COMPARING MAPS **//

	// Maps cannot be compared using == operator. A map can be compared only to nil.
	a := map[string]string{"A": "X"}
	b := map[string]string{"B": "X"}

	// fmt.Println(a == b) // error -> invalid operation: a == b (map can only be compared to nil)

	// to compare 2 maps that have the keys and values of type string
	// we get a string representation of the maps and compare those strings.

	// getting a string representation of maps called a and b
	s1 := fmt.Sprintf("%s", a)
	s2 := fmt.Sprintf("%s", b)

	if s1 == s2 {
		fmt.Println("Maps are equal")
	} else {
		fmt.Println("Maps are not equal")
	}

	//** CLONING A MAP **//

	// When creating a map variable Go creates a pointer to a map header value in memory.
	// The key: value pairs of the map are not stored directly into the map.
	// They are stored in memory at the address referenced by the map header.

	friends := map[string]int{"Dan": 40, "Maria": 35}

	// neighbors is not a copy of friends.
	// both maps reference the same data structure in memory
	neighbors := friends

	// modifying friends AND neighbors
	friends["Dan"] = 30

	fmt.Println(neighbors) // -> map[Dan:30 Maria:35]

	// How to clone a map?
	// 1. initialize a new map
	colleagues := make(map[string]int)

	// colleagues = friends // -> ERROR, illegal with maps!

	// 2. use a for loop to copy each element into the new map
	for k, v1 := range friends {
		colleagues[k] = v1
	}

	// colleagues and friends are different maps in memory!

}
