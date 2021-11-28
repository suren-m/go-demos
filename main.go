package main

import (
	"fmt"
	"time"

	"rsc.io/quote"
)

func main() {
	fmt.Println("starting")

	//	var wg sync.WaitGroup
	c1 := make(chan string)
	c2 := make(chan string)
	//wg.Add(2)

	// msg := <-c
	// fmt.Println(msg)

	go network_call_1(c1)
	go network_call_2(c2)

	// for msg := range c1 {
	// 	fmt.Println(msg)
	// }

	// for msg := range c2 {
	// 	fmt.Println(msg)
	// }

	for {
		select {
		case msg1, open := <-c1:
			if !open {
				c1 = nil
			}
			fmt.Println(msg1)
		case msg2, open := <-c2:
			if !open {
				c2 = nil
			}
			fmt.Println(msg2)
		}

		if c1 == nil && c2 == nil {
			break
		}
	}

	//	wg.Wait()
	fmt.Println("done")

	//print_json()
	//pointer_printer()
}

func network_call_1(c chan<- string) {
	defer close(c)
	for i := 1; i <= 5; i++ {
		time.Sleep(2 * time.Second)
		msg := quote.Go()
		c <- msg
	}
}

func network_call_2(c chan<- string) {
	defer close(c)
	for i := 1; i <= 5; i++ {
		time.Sleep(1 * time.Second)
		msg := quote.Opt()
		c <- msg
	}
}
