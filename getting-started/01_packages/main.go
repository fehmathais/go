package main

import (
	"fmt"
	"module/utils"
)

// the command: go build
// will compile everything in the package in a single and runable binary
func main() {
	fmt.Println("Writting from main package...")
	utils.PrintText("Writting from main utils...")
}
