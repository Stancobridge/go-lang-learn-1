package main

import (
	"fmt"
	"log"
	"modules/finance"
)

func main() {
	log.SetPrefix("greetings: Hi")
	fmt.Println(finance.WorkWithSlice())
	message, err := finance.CalculateTime("Stanley")

	if err != nil {

		log.Fatal(err)
	}
	fmt.Println(message)

}
