package main

import "fmt"

func main() {
	var firstName string

	firstName = "Felix"
	fmt.Println(firstName)

	lastName := "Ruby"
	fmt.Println(lastName)

	// multiple variable
	var (
		age   = 22
		title = "Fullstack Developer"
	)

	fmt.Println(age)
	fmt.Println(title)

	city, isMarried := "Bandung", false
	fmt.Println(city)
	fmt.Println(isMarried)
}
