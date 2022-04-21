package main

import "fmt"

func main() {
	var names [2]string

	names[0] = "Felix"
	names[1] = "Ruby"

	fmt.Println(names[0])
	fmt.Println(names[1])

	var score = [3]int{
		70,
		80,
		90,
	}

	fmt.Println(score)

	var emptyArr [10]string
	fmt.Println(len(emptyArr)) // print array length

	values := [...]int{
		10,
		20,
		30,
		40,
	}

	fmt.Println(values)
	fmt.Println(len(values))
	values[0] = 100
	fmt.Println(values)
}
