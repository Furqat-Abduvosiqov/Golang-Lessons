//--Summary:
//  Create a program that can perform dice rolls using various configurations
//  of number of dice, number of rolls, and number of sides on the dice.
//  The program should report the results as detailed in the requirements.
//
//--Requirements:
//* Print the sum of the dice roll
//* Print additional information in these circumstances:
//  - "Snake eyes": when the total roll is 2, and total dice is 2
//  - "Lucky 7": when the total roll is 7
//  - "Even": when the total roll is even
//  - "Odd": when the total roll is odd
//* The program must use variables to configure:
//  - number of times to roll the dice
//  - number of dice used in the rolls
//  - number of sides of all the dice (6-sided, 10-sided, etc determined
//    with a variable). All dice must use the same variable for number
//    of sides, so they all have the same number of sides.
//
//--Notes:
//* Use packages from the standard library to complete the project
//* Try using different values for your variables to check the results

package main

import "math/rand"

func main() {
	// Configuration variables
	numRolls := 5
	numDice := 2
	numSides := 6
	// Roll the dice and store results
	rolls := rollDice(numRolls, numDice, numSides)

	printSumOfRolls(rolls)
}

func printSumOfRolls(rolls []int) {
	sum := 0
	for _, roll := range rolls {
		sum += roll
	}
	println("Sum of rolls:", sum)
}

func rollDice(numRolls, numDice, numSides int) []int {
	rolls := make([]int, numRolls)
	for i := 0; i < numRolls; i++ {
		rolls[i] = rollSingle(numDice, numSides)
	}	
	return rolls
}

func rollSingle(numDice, numSides int) int {
	sum := 0	
	for i := 0; i < numDice; i++ {
		sum += rollDie(numSides)
	}		
	return sum
}

func rollDie(numSides int) int {
	return rand.Intn(numSides) + 1
}

