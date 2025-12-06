package main

import "fmt"

const age = 30

var ice string = "cold" // valid
// hot := "summer"  invalid

func main() {

	const name string = "golang"
	// const weight float32

	// weight = 33.3

	// name = "java"  incorrect as we can't reassign the const

	fmt.Println(age)

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(port)
}