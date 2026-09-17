package main

import "fmt"

func main() {
	// Constants - value never changes
	const maxUpload = 2

	const canUpload bool = true
	// canUpload = false  /* Cannot change. It will throw error */

	fmt.Println("Maximum file upload size is", maxUpload)
	fmt.Println("Can upload files?", canUpload)

}