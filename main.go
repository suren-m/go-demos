package main

import (
	"fmt"

	"rsc.io/quote"

	"services"
)

func main() {

	fmt.Println("hello")

	// calling an exported func from external package
	fmt.Println(quote.Go())

	services.Print_Json()

}
