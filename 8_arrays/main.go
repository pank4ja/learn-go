package main

import "fmt"

func main() {

	var nums [4]int

	nums[0] = 1

	fmt.Println(nums[0])
	fmt.Println(nums)

	// array length
	fmt.Println(len(nums))

	var vals [4]bool

	fmt.Println(vals)

	var str [5]string
	fmt.Println(str)

	str[0]=  "golang"
	fmt.Println(str)


	// to declare 
	n := [3]int {1,2,3}
	fmt.Println(n)


	// 2d array
	m := [2][2]int{{1,2},{3,4}}
	fmt.Println(m)
}