package main

import (
	"fmt"
	"strconv"
)

func main() {
	var x = 3
	var y = 3.1
	x = x * int(y)
	fmt.Println(x)
	fmt.Printf("%v %T\n", y, y)

	fmt.Printf("%T\n", int(y))
	x = int(float64(x) * y)
	fmt.Printf("%T\n", x)

	s := string(99)
	fmt.Println(s)

	//s1:= string(44.2)

	var myStr = fmt.Sprintf("%f", 44.2)
	fmt.Println(myStr)

	var myStr1 = fmt.Sprintf("%d", 343434)
	fmt.Println(myStr1)

	fmt.Println(string(43434))

	s1 := "3.123"
	fmt.Printf("%T\n", s1)

	var f1, err = strconv.ParseFloat(s1, 64)
	_ = err
	fmt.Println(f1 * 3.4)

	i, err := strconv.Atoi("-50")
	_ = err
	s2 := strconv.Itoa(20)
	fmt.Printf("i type is %T, i value is %v", i, i)
	fmt.Printf("s2 type is %T, s2 value is %q", s2, s2)

}
