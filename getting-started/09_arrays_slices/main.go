package main

import "fmt"

func main() {
	a := [5]int{1, 2, 3, 4, 5}   // we can declare fixed array in this way
	b := [...]int{1, 2, 3, 4, 5} // in this case, go will predefine the size/cap of the array
	c := b
	c[0] = 100 // this will not affect the original array

	fmt.Printf("Arrays are fixed in size and cap: %v, %v\n", len(a), cap(a))
	fmt.Println("A: ", a)
	fmt.Println("B: ", b)
	fmt.Printf("C:  %v", c)
	fmt.Printf(" when we modify the first element of C, it will not affect the original array\n")

	fmt.Println("---------------------------")
	fmt.Println("Slicing fixed arrays...")
	// we can slice an fixed array
	fmt.Println("C[:] ", b[:])     // [1,2,3,4,5]
	fmt.Println("C[1:] ", b[1:])   // [2,3,4,5]
	fmt.Println("C[:3] ", b[:3])   // [1,2,3]
	fmt.Println("C[2:4] ", b[2:4]) // [3,4]

	fmt.Println("---------------------------")
	fmt.Println("Slices:")
	d := []int{1, 2, 3, 4, 5} // this is known as a slice in go
	e := d                    // slices are references (pointers) to the original array
	e[0] = 100                // this will affect the original array
	fmt.Println("D: ", d)
	fmt.Println("E: ", e)

	fmt.Printf("Arrays are not fixed in size and cap: %v, %v\n", len(e), cap(e))
	e = append(e, []int{6, 7, 8}...) // this will affect the original array
	fmt.Println("\nE: ", e)
	fmt.Printf("Now their size and cap are: %v, %v\n", len(e), cap(e))

	// but when we append a slice we don't modify the original based array
	fmt.Println("D: ", d)
	fmt.Println("E: ", e)
	fmt.Printf("Now their size and cap of d are: %v, %v\n", len(d), cap(d))
	fmt.Printf("Now their size and cap of e are: %v, %v\n", len(e), cap(e))

	fmt.Println("---------------------------")
	fmt.Println("Slicing dynamic arrays...")
	// we can also slice a dynamic array
	fmt.Println("D[2:4] ", d[2:4])
	fmt.Println("D[:3] ", d[:3])
	fmt.Println("D[2:] ", d[2:])
	fmt.Println("D[:] ", d[:])
}
