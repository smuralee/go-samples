package main

import "fmt"

func main() {
	fmt.Println("Hello World!!")

	fmt.Println("Sum : ", 2+6)
	fmt.Println("Divide : ", 2/5)

	var isValid = true
	var isInvalid = false
	fmt.Println("Boolean - true & false : ", isValid && isInvalid)
	fmt.Println("Boolean - true or false : ", isValid || isInvalid)
	fmt.Println("Boolean - negation true : ", !isValid)

	const n = 500
	const d = 3e20/n
	fmt.Println(d)
}
