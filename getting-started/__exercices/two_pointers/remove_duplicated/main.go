package main

import "fmt"

func main() {
	array := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	var j int = 0

	for i := 0; i < len(array); i++ {
		if array[i] != array[j] {
			j++
			array[j] = array[i]
		}
	}

	fmt.Println(array[:j+1])
}
