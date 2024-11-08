package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// Since if and switch accept an initialization statement, it's common to see one used to set up a local variable.
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	}
}
