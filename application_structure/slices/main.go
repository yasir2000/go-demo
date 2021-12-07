package main

import (
	"fmt"
	"unsafe"
)

func main() {
	// declaring a string slice, by default is initialized with nil or uninitialized
	var cities []string

	fmt.Println("cities is equal to nil: ", cities == nil) // -> cities is equal to nil:  true
	fmt.Printf("cities: %#v\n", cities)                    // -> cities: []string(nil)

	// we can not assign elements to nil slice:
	// cities[0] = "Paris" // -> runtime error

	// declaring a slice using a slice literal
	numbers := []int{2, 3, 4, 5} // on the right hand-side of the equal sign is a slice literal
	fmt.Println(numbers)         // => [2 3 4 5]

	// creating a slice using the make() built-in function
	// creating a slice with 2 int elements initialized with zero.
	nums := make([]int, 2) // the same as []int{0, 0}.
	fmt.Println(nums)      // => [0 0]

	// declaring a slice using a slice literal
	type names []string
	friends := names{"Dan", "Maria"}
	fmt.Println(friends)

	// getting an element from the slice
	x := numbers[0]
	fmt.Println("x is", x) // => x is 2

	// modifying an element of the slice
	numbers[1] = 200
	fmt.Printf("%#v\n", numbers) // => []int{2, 200, 4, 5}

	// iterating over a slice
	for index, value := range numbers {
		fmt.Printf("index: %v, value: %v\n", index, value)
	}

	numbers = append(numbers, 192, 87, 45, 11)
	fmt.Println(numbers)

	n := []int{87, 92}
	numbers = append(numbers, n...)
	fmt.Println(numbers)

	src := []int{14, 21, 345}
	dst := make([]int, 2)
	nn := copy(dst, src)
	fmt.Println(src, dst, nn)

	a := [5]int{1, 3, 4, 5, 6}
	//a[start:stop]
	b := a[1:3]
	fmt.Printf("%v, %T\n", b, b)

	s1 := []int{1, 2, 3, 4, 5, 6}
	s2 := s1[1:3]
	fmt.Println(s2)

	fmt.Println(s1[2:])
	fmt.Println(s1[:3])
	fmt.Println(s1[:])

	//s1 = append(s1[2:4], 30)
	//fmt.Println(s1)

	//s1 = append(s1[:4], 200)
	//fmt.Println(s1)

	s1 = append(s1[3:], 128)
	fmt.Println(s1)

	s3 := []int{20, 10, 30, 88, 876}
	s4, s5 := s3[0:2], s3[1:3]

	s4[1] = 876
	fmt.Println(s3)
	fmt.Println(s5)

	arr1 := [4]int{93, 45, 21, 45}
	slice1, slice3 := arr1[0:2], arr1[1:3]
	arr1[1] = 2
	fmt.Println(arr1, slice1, slice3)

	cars := []string{"Ford", "Honda", "Audi", "Range Rover"}
	newCars := []string{}

	newCars = append(newCars, cars[0:2]...)

	cars[0] = "Nissan"
	fmt.Println(cars, newCars)

	s8 := []int{10, 20, 30, 40, 50}

	newSlice := s8[0:3]
	fmt.Println(len(newSlice), cap(newSlice))

	newSlice = s1[2:5]
	fmt.Println(len(newSlice), cap(newSlice))

	fmt.Printf("%p\n", &s8)

	fmt.Printf("%p %p \n", &s8, &newSlice)
	newSlice[0] = 1000
	fmt.Println("s1:", s1)

	j := [5]int{1, 2, 3, 4, 5}
	s := []int{1, 2, 3, 4, 5}

	fmt.Printf("array's size in bytes : %d \n", unsafe.Sizeof(j))
	fmt.Printf("array's size in bytes : %d \n", unsafe.Sizeof(s))

	numbers1 := []int{2, 3}

	// append() returns a new slice after appending a value to its end
	numbers1 = append(numbers1, 10)
	fmt.Println(numbers) //-> [2 3 10]

	// appending more elements at once
	numbers1 = append(numbers1, 20, 30, 40)
	fmt.Println(numbers1) //-> [2 3 10 20 30 40]

	// appending all elements of a slice to another slice
	n1 := []int{100, 200, 300}
	numbers1 = append(numbers1, n1...) // ... is the ellipsis operator
	fmt.Println(numbers1)              // -> [2 3 10 20 30 40 100 200 300]

	//** Slice's Length and Capacity **//

	nums1 := []int{1}
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums1), cap(nums1)) // Length: 1, Capacity: 1

	nums1 = append(nums1, 2)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums1), cap(nums1)) // Length: 2, Capacity: 2

	nums1 = append(nums1, 3)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums1), cap(nums1)) // Length: 3, Capacity: 4
	// the capacity of the new backing array is now larger than the length
	// to avoid creating a new backing array when the next append() is called.

	nums1 = append(nums1, 4, 5)
	fmt.Printf("Length: %d, Capacity: %d \n", len(nums1), cap(nums1)) // Length: 5, Capacity: 8

	// copy() function copies elements into a destination slice from a source slice and returns the number of elements copied.
	// if the slices don't have the same no of elements, it copies the minimum of length of the slices
	src1 := []int{10, 20, 30}
	dst1 := make([]int, len(src1))
	nn1 := copy(dst1, src1)
	fmt.Println(src1, dst1, nn1) // => [10 20 30] [10 20 30] 3

}
