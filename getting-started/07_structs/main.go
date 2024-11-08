package main

import (
	"fmt"
)

type Person struct {
	name   string
	age    int
	father *Person
}

type Employee struct {
	Person
	salary float32
}

func main() {
	var e Employee
	e.name = "Felipe"
	e.age = 32
	e.salary = 10000.00
	e.father = &Person{"Ordilafo", 50, nil}
	fmt.Printf("Actually %s is %d years old and his salary is %.2f\n", e.name, e.age, e.salary)
	fmt.Printf("The name of %s father's is %s\n", e.name, e.father.name)
}
