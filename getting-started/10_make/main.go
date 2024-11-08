package main

import "fmt"

func main() {
	// make is used to create arrays, slices, maps, and channels only
	v := make([]int, 10, 100) // this will create an array of 10 zeros with a capacity of 100
	z := make([]int, 10)      // this way will define capacity automatically

	fmt.Println(v, len(v), cap(v))
	fmt.Println(z, len(z), cap(z))

	fmt.Println("---------------------------")
	// make its helpful to create a 2D array, eg: images
	xSize := 4 // an example of image columns (width)
	ySize := 4 // an example of image rows (height)

	// unic allocation is more efficient when we have a fixed size of lines
	pa := make([][]uint8, ySize)
	for i := range pa {
		pa[i] = make([]uint8, xSize) // example of unic allocation
		fmt.Println(pa[i])
	}

	fmt.Println("---------------------------")

	// dinamically allocated, is helpful when we need to resize the array
	pb := make([][]uint8, ySize)
	pixels := make([]uint8, xSize*ySize)
	for i := range pb {
		pb[i], pixels = pixels[:xSize], pixels[xSize:]
		fmt.Println(pa[i])
	}
}
