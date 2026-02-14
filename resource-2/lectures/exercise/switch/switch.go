//--Summary:
//  Create a program to display a classification based on age.
//
//--Requirements:
//* Use a `switch` statement to print the following:
//  - "newborn" when age is 0
//  - "toddler" when age is 1, 2, or 3
//  - "child" when age is 4 through 12
//  - "teenager" when age is 13 through 17
//  - "adult" when age is 18+

package main

func main() {

	switch age := 14; {
	case age == 0:
		println("newborn")
	case age == 1 || age == 2 || age == 3:
		println("toddler")
	case age >= 4 && age <= 12:
		println("child")
	case age >= 13 && age <= 17:
		println("teenager")
	default:
		println("adult")
	}
}
