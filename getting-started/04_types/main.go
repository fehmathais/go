package main

import (
	"errors"
	"fmt"
)

func main() {
	var a bool = true // Boolean

	// if we don't specify the bit size, the compiler will infer it according to the system architecture 32 or 64bits.
	// we can also declare int8, int16, int32, int64
	var b int = 5       // Integer.
	var b2 int = -10000 // An integer can also be negative
	var b3 uint = 10000 // Unsigned values can't be negative

	// runes are aliases for int32
	// they generaly are used to represent unicode characters of ASCII or UTF-8
	// in other words, an array of characters can be represented as an array of runes
	var c rune = 'a' // rune is an alias for int32
	var c2 rune = '😀'

	// we can also declare float32 and float64
	var d float32 = 3.14 // Floating point number.

	// a conventional string
	var e string = "Hi!" // String

	// go don't have char type
	// var f := 'a' // Character

	// bytes are like an alias for uint8
	var g byte = 'a'
	var g2 byte = 123

	// errors are types in go
	// if the value of the error is nil, it means that there was no error
	// nil value represents the state of 'zero value' (an empty value)
	var err error = errors.New("Something went wrong")

	fmt.Println("Boolean: 			", a)
	fmt.Println("Integer: 			", b)
	fmt.Println("Negative Integer: 		", b2)
	fmt.Println("Unsigned Integer: 		", b3)
	fmt.Println("Runes for ASCII:		", c)
	fmt.Println("Runes for emoji: 		", c2)
	fmt.Println("Float:   			", d)
	fmt.Println("String:  			", e)
	fmt.Println("Bytes for ASCII:		", g)
	fmt.Println("Bytes for numbers:		", g2)
	fmt.Println("Error:				", err)
}
