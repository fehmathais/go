package utils

// if a function is written with a capital letter, it's exported
// if a function is written with a lowercase letter, it's not exported and can only be used in the same package
func PrintText(text string) {
	print(text)
}
