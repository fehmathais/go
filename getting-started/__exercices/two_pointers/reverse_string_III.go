package main

import (
	"fmt"
)

func main() {
	var s string = "Estou a gostar de programar em Go!"
	words := []rune(s)
	start := 0

	fmt.Println(string(words))

	for i := 0; i <= len(words); i++ {
		if i == len(words) || words[i] == ' ' {
			reverse(words, start, i-1)
			start = i + 1
		}
	}

	fmt.Println(string(words))
}

func reverse(word []rune, start int, end int) {
	for start < end {
		word[start], word[end] = word[end], word[start]
		start++
		end--
	}
}
