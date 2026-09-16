package main

import "fmt"

func main() {
	// Basic printing
	var age int
	var name string = "Parmesh"
	age = 22
	var radius float64 = 5.5

	isStudent := true

	fmt.Println("Student", isStudent)
	fmt.Println("My radius is", radius)
	fmt.Println("My name is", name)
	fmt.Println("My age is", age)

	// Sum with type infer
	number1 := 10
	number2 := 20

	var number3 = number1 + number2
	fmt.Println("Sum of", number1, "and", number2, "is", number3)
}
