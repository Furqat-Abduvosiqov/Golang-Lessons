//--Summary:
//  Use functions to perform basic operations and print some
//  information to the terminal.
//
//--Requirements:
//* Write a function that accepts a person's name as a function
//  parameter and displays a greeting to that person.
//* Write a function that returns any message, and call it from within
//  fmt.Println()
//* Write a function to add 3 numbers together, supplied as arguments, and
//  return the answer
//* Write a function that returns any number
//* Write a function that returns any two numbers
//* Add three numbers together using any combination of the existing functions.
//  * Print the result
//* Call every function at least once

package main

func main() {
	greetPerson("Furqat")
	println(returnAnyMessage())
	println("Sum of 1, 2, 3 is", addThreeNumbers(1, 2, 3))
	println("Random number from function", returnAnyNumber())
	num1, num2 := returnAnyTwoNumbers()
	println("Two numbers returned from function:", num1, num2)
	sum := addThreeNumbers(num1, num2, returnAnyNumber())
	println("Sum of", num1, num2, "and", returnAnyNumber(), "is", sum)
}

func greetPerson(personName string) {
	println("Hello", personName)
}

func returnAnyMessage() string {
	return "This is random message"
}

func addThreeNumbers(num1, num2, num3 int) int {
	return num1 + num2 + num3
}

func returnAnyNumber() int {
	return 42
}

func returnAnyTwoNumbers() (int, int) {
	return 10, 20
}
