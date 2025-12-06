package main

import "fmt"

func main() {

	var name string = "value"

	// every variable must be used if declared

	fmt.Println(name)

	var name2 = "golang"  // automatic infer the type if type is not declared
	fmt.Println(name2)

	var isAdult bool= true
	fmt.Println(isAdult)


	var age int = 30
	fmt.Println(age)

	// shortcut to declare  when we have to declare the variable and also assign the value to it also infer type automatically
	naam := "goland"
	fmt.Println(naam)

	// when we declare the variable value later we have to use type (just for defining)
	var name4 string

	name4 = "late declare"
	fmt.Println(name4)

	var price float32 = 33.2
	// var price = 50.5
	// price := 50.3
	fmt.Println(price)
}