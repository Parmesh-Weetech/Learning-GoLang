package main

import "fmt"

func main() {
	n := 10
	total := 0
	for i := 0; i <= n; i++ {
		total += i
	}

	fmt.Println(total)
}