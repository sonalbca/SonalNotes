package main

import "fmt"

/*
Simple function with return value in Golang
In this example, the add() function takes input of two integer numbers and returns an integer value with a name of total.

Note the return statement is required when a return value is declared as part of the function's signature.
*/

// Function with int as return type
func Add(x int, y int) int {
	total := 0
	total = x + y
	return total
}

func main() {
	// Accepting return value in varaible
	sum := Add(20, 30)
	fmt.Println(sum)
}

/*
    The types of input and return value must match with function signature.
	If we will modify the above program and pass some string value in argument then program will throw an exception
    "cannot use "test" (type string) as type int in argument to add".
*/
