package main

import (
	"fmt"
	"slices"
)

func main() {

	// uninitalized slice in nil
	// var nums []int

	// fmt.Println(nums == nil)
	// fmt.Println(len(num))

	var nums = make([]int, 2, 5)  // 5 is inital capacity
	fmt.Println(cap(nums))  // capacity  maximum no. of elements can be fixed ( but it can be rezied dynamic)
	// fmt.Println(nums == nil)
	nums = append(nums, 1)
	nums = append(nums, 2)
	nums = append(nums, 3)
	nums = append(nums, 4)
	nums = append(nums, 5)
	nums = append(nums, 6)
	nums = append(nums, 7)
	nums = append(nums, 8)
	nums = append(nums, 9)
	nums = append(nums, 10)
	nums = append(nums, 11)
	fmt.Println(nums)
	fmt.Println(cap(nums))

	var emp = make([]int, 0, 5)
	fmt.Println(emp)



	t := []int {} // create an empty slice
	t = append(t, 4)
	fmt.Println(t)
	fmt.Println(cap(t))
	fmt.Println(len(t))


	var a = make([]int, 2,2)
	a[0] = 1
	a[1] = 4
	// a[2] = 3

	fmt.Println(cap(a))
	fmt.Println(a)

	var test = make([]int, 0, 5)
	test = append(test, 2)
	var test2 = make([]int, len(test))


	copy(test2, test)

	fmt.Println(test, test2)


	// slice operator

	var sli = []int {1,2,3,4,5}
	fmt.Println(sli[0:2])
	fmt.Println(sli[:1])
	fmt.Println(sli[1:])
	fmt.Println(sli[:])


	var n1 = []int {1,2}
	var n2 = []int {1,2}

	fmt.Println(slices.Equal(n1,n2))


	var darr = [][]int {{1,2,3},{4,5,6}}
	fmt.Println(darr)
}