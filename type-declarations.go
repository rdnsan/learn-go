package main

import "fmt"

func main() {
	type NIK string
	type Married bool

	var noKtp NIK = "32078978794564646"
	var isMarried Married = false

	fmt.Println("NIK :", noKtp)
	fmt.Println("Is Married:", isMarried)
}
