package main

import "fmt"

func main() {

	// declared in the outer scope of the main function
	var a int

	// You can change this to "internal" or "external" to test both cases
	source := "external"

	if source == "internal" {
		// This 'a' is a new variable, local to the if block
		// It shadows the outer 'a'
		// Any reference to 'a' inside this block refers to this local one
		// So 'a' here is 10, but the outer 'a' remains unchanged (still 0)
		a := 10
		fmt.Println("value of a inside if is ", a)
	} else {
		// Here, instead of declaring a new variable with ':=',
		// we're assigning to the outer 'a' declared above
		a = 15
		fmt.Println("value of a inside else is ", a)
	}
	fmt.Println(a) // prints 0 if source is internal or else prints 15
}
