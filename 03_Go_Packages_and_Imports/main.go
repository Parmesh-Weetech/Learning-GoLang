package main

import (
	"fmt"
	"math"
)

func main() {
	var radius float64 = 5.5
	area := math.Pi * math.Pow(radius, 2)
	
	fmt.Println("Area of circle is", area)
}
