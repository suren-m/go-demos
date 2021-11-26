package main

import (
	"log"
)

func pointer_printer() {
	var greeting string
	greeting = "hello"

	audience := "world"
	log.Println(greeting, audience)

	ptr := &audience
	log.Println("addr:", ptr, "data:", *ptr)

	*ptr = "galaxy"
	log.Println(greeting, audience, *ptr)

	doubleptr := &ptr
	**doubleptr = "universe"

	log.Println(greeting, audience, *ptr, **doubleptr)

	local_changer(audience)
	log.Println(greeting, audience, *ptr, **doubleptr)

	changer(&audience)
	log.Println(greeting, audience, *ptr, **doubleptr)
}

func local_changer(audience string) {
	audience = "world again"
	log.Println("locally changing to", audience)
}

func changer(audience *string) {
	*audience = "world again"
	log.Println("changing to", *audience)
}
