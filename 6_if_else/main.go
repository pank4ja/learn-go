package main

import "fmt"

func main() {

	var age int = 14

	if age >= 18 {
		fmt.Println("eligible")
	} else {
		fmt.Println("not eligible")
	}


	price := 20

	if price >= 18 {
		fmt.Println("high")
	} else if price >=10{
		fmt.Println("moderate")
	} else {
		fmt.Println("low")
	}


	var role = "admin"
	var haspermission = true

	if role == "a2dmin" || haspermission {
		fmt.Println("yes matched")
	}

	if role == "a2dmin" && haspermission {
		fmt.Println("testing")
	}

	if age := 15; age >= 18 {
		fmt.Println("valid", age)
	} else if age >= 12{
		fmt.Println("teenage", age)
	}
}