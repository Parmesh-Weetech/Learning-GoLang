package main

import "fmt"

func main() {
	n := 2

	switch n {
	case 1:
		fmt.Println("One")
	case 2:
		fmt.Println("Two")
	case 3:
		fmt.Println("Three")
	default:
		fmt.Println("Number is greater than 3")
	}
}
