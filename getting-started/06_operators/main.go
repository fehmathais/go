package main

import "fmt"

func main() {
	var soma = 2 + 2
	var subtracao = 2 - 2
	var multiplicacao = 2 * 2
	var divisao = 2 / 2
	var resto = 2 % 2

	fmt.Println(soma, subtracao, multiplicacao, divisao, resto)

	// it's not possible to do operations in go with different types
	// var x int = 10
	// var y string = "10"
	// var z = x + y
	// fmt.Println(z)

	// even for the same type but different archs
	// var x int = 10
	// var y int8 = 10
	// var z = x + y
	// fmt.Println(z)

	// we can use the unary operator to invert the value of a variable
	var a int = 10
	a++
	a--
	a += 10
	a -= 10
	a *= 10
	a /= 10
	a %= 10
	fmt.Println(a)

	// attribution operator
	var c int = 10
	d := 20
	c += 10
	fmt.Println(c)
	fmt.Println(d)

	// relational operators
	var e int = 10
	var f int = 20
	fmt.Println(e < f)
	fmt.Println(e > f)
	fmt.Println(e <= f)
	fmt.Println(e >= f)
	fmt.Println(e == f)
	fmt.Println(e != f)

	// logical operators
	var g bool = true
	var h bool = false
	fmt.Println(g && h)
	fmt.Println(g || h)
	fmt.Println(!g)

	// bitwise operators
	var i int = 1
	var j int = 2

	fmt.Printf("bitwise &: %b \n", i&j) // Operação AND (i & j): 0001 & 0010 = 0000
	// Saída: 0 (em binário 0000) - os bits de i e j são comparados, e apenas posições com 1 em ambos resultarão em 1.

	fmt.Printf("bitwise |: %b \n", i|j) // Operação OR (i | j): 0001 | 0010 = 0011
	// Saída: 11 (em binário 0011) - compara os bits de i e j, e as posições com pelo menos um 1 resultam em 1.

	fmt.Printf("bitwise ^: %b \n", i^j) // Operação XOR (i ^ j): 0001 ^ 0010 = 0011
	// Saída: 11 (em binário 0011) - compara os bits de i e j, e apenas as posições com bits diferentes resultam em 1.

	fmt.Printf("bitwise <<: %b \n", i<<j) // Operação SHIFT LEFT (i << j): 0001 << 2 = 0100
	// Saída: 100 (em binário 0100) - desloca os bits de i (0001) duas posições à esquerda, resultando em 0100 (decimal 4).

	fmt.Printf("bitwise >>: %b \n", j>>i) // Operação SHIFT RIGHT (j >> i): 0010 >> 1 = 0001
	// Saída: 1 (em binário 0001) - desloca os bits de j (0010) uma posição à direita, resultando em 0001 (decimal 1).
}
