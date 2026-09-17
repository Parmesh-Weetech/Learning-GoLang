package main

import "fmt"

func main() {
	item := 3
	pricePerItem := 10

	if total := item * pricePerItem; total > 20 {
		fmt.Println("More than 30")
	} else {
		fmt.Println("Less than 30")
	}
}