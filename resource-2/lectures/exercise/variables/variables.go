//Summary:
//  Print basic information to the terminal using various variable
//  creation techniques. The information may be printed using any
//  formatting you like.
//
//Requirements:
//* Store your favorite color in a variable using the `var` keyword
//* Store your birth year and age (in years) in two variables using
//  compound assignment
//* Store your first & last initials in two variables using block assignment
//* Declare (but don't assign!) a variable for your age (in days),
//  then assign it on the next line by multiplying 365 with the age
// 	variable created earlier
//
//Notes:
//* Use fmt.Println() to print out information
//* Basic math operations are:
//    Subtraction    -
// 	  Addition       +
// 	  Multiplication *
// 	  Division       /

package main

func main() {
	var favoriteColor string = "black"
	birthYear, age := 1995, 31

	var (
		firstInitial string = "A"
		lastInitial  string = "F"
	)
	var ageInDays int
	ageInDays = age * 365

	println("My favorite color is", favoriteColor)
	println("I was born in", birthYear, "and I am", age, "years old.")
	println("My initials are", firstInitial, lastInitial)
	println("I am approximately", ageInDays, "days old.")
	println("I must to pray to the Allah", ageInDays * 5 , "times in my life.")
}

