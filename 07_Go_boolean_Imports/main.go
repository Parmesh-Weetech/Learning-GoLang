package main

import (
	"fmt"
)

func main() {
	// Simple boolean flag
	var isLoggedIn bool = true
	fmt.Println("Logged In: ", isLoggedIn)

	isLoggedIn = false
	fmt.Println("Logged In: ", isLoggedIn)

	// Logical operators
	isLoggedIn = true
	isAdmin := false
	hasPermission := true
	isLoggedIn = isLoggedIn && isAdmin
	fmt.Println("Logged In After logical AND: ", isLoggedIn)

	hasAccess := isLoggedIn || hasPermission
	fmt.Println("Has Access After logical OR: ", hasAccess)
}
