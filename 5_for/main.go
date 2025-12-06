package main

import "fmt"

func main() {

	// no while in go
	i := 1

	// while loop
	for i <= 3 {
		fmt.Println(i)
		i = i+1
	}

	// this is infinite loop
	// for {
	// 	println(1)
	// }


	// classic for loop
	for i := 0 ; i<= 3 ; i++ {
		// break

		if i== 2 {
			continue
		}

		// continue skips that particular interation
		fmt.Println(i)
	}


	for i:= range 3 {
		fmt.Println(i)
	}
}