package main

import (
	"errors"
	"fmt"
)

var d = func(a int, b int) int {
	return a + b
}

// a function can return multiple values which is very helpful for error handling
func f(a int, b int) (int, error) {
	if (a < 0) || (b < 0) {
		err := errors.New("Only positive numbers are allowed")
		return 0, err
	}

	return a + b, nil
}

func main() {
	var a int = -10
	var b int = 20

	// we can declaed a function inside another function and assign it to a variable
	var c = func() string {
		return fmt.Sprintf("%d", a) // it's possible to access a variable declared outside the function
	}

	var result, err = f(a, b)

	fmt.Println("Value of a:		", c())
	fmt.Println("Sum of a and b:		", d(a, b))

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("Sum with err handling:	", result)
	}
}
