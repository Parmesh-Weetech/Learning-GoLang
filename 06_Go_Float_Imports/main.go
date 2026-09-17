package main

import (
	"fmt"
	"math"
)

func main() {

	var radius float64 = 5.5

	circumference := 2 * math.Pi * radius
	fmt.Printf("Circumference of the circle: %.2f\n", circumference)
}
