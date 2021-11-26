package main

import (
	"fmt"

	"rsc.io/quote"
)

func main() {
	var greeting string
	greeting = "hello"
	audience := "world"
	fmt.Println(greeting, audience)

	// calling an exported func from external package
	fmt.Println(quote.Go())

	print_json()
}
