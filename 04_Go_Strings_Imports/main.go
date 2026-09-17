package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.ToUpper("hello, world!"))

	var contains = strings.Contains("hello, world!", "world")

	fmt.Println("Does world is inside hello world! word? ", contains)
}
