package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {

	fmt.Println("hello")

	// calling an exported func from external package
	fmt.Println(quote.Go())

	print_json()

	pointer_printer()
}
