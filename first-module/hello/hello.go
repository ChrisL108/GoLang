package main

import (
	"fmt"
	"log"

	"example.com/greetings"
)

func main() {

	log.SetPrefix("greetings: ")
	log.SetFlags(0)

	message, err := greetings.Hello("Christopher")
	if err != nil {
		log.Fatal(err)
	}
	myNum := test(4)
	fmt.Println(myNum)
	fmt.Println(message)
}

func test(num int) int {
	return num
}
