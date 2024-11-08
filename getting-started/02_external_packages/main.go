package main

import (
	"fmt"

	"github.com/badoux/checkmail"
)

func main() {
	var input string
	fmt.Print("Write your email: ")
	fmt.Scanln(&input)
	err := checkmail.ValidateFormat(input)

	if err != nil {
		fmt.Println("Email is not valid")
	} else {
		fmt.Println("Email is valid")
	}
}
