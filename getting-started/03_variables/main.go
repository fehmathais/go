package main

import "fmt"

func main() {
	// we can declare type of a variable or we can infer it
	var a string = "Declared with type"
	fmt.Println(a)

	b := "Declared without type"
	// as similar to other languages, an string is an "array of characters"
	fmt.Println(b[1]) // 101 is the ascii code of "e" in decimal
	fmt.Println(&b)   // address of the string

	// print the binary value of b
	for _, c := range b {
		fmt.Printf("%b ", c) // %b is the binary format specifier
		// 1000100 1100101 1100011 1101100 1100001 1110010 1100101 1100100 100000 1110111 1101001 1110100 1101000 1101111 1110101 1110100 100000 1110100 1111001 1110000 1100101 %
	}

	fmt.Println("")

	// we also can declare variables like this
	var (
		c string = "Variable C"
		d string = "Variable D"
	)

	// and we also can print them in a single instruction
	fmt.Println(c, d)

	// we can invert the value of a variable
	c, d = d, c
	fmt.Println(c, d)

	// constants, as in other languages we can't change its value once it's declared
	const e = "Constant E"
	fmt.Println(e)
	// e := "New value" // it even doesn't compile
}
