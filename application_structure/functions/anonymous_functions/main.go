package main

import "fmt"

func increment(x int) func() int {
	return func() int {
		x++
		return x
	}
}
func twoSum(nums []int, target int) []int {
	memo := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		j, found := memo[target-nums[i]]
		if found {
			return []int{j, i}
		}
		memo[nums[i]] = i
	}
	return []int{}
}
func main() {
	func(msg string) {
		fmt.Println(msg)
	}("I'm an anonymous function!")

	a := increment(10) //first class functions support
	fmt.Printf("%T\n", a)
	a()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	s := twoSum(8, 3)
	fmt.Println(s)

}
