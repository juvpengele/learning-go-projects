package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println("Welcome to the Number Guessing Game!")
	fmt.Println("I'm thinking of a number between 1 and 10.")



	// Create a random number between 1 and 10 but it should be random every time
	var guessingNumber = rand.Intn(10)

	var userGuess int

	fmt.Println("Can you guess what it is?")
	fmt.Scan(&userGuess)
	
	for userGuess != guessingNumber {
		if userGuess < guessingNumber {
			fmt.Println("Too low! Try again.")
		} else {
			fmt.Println("Too high! Try again.")
		}
		fmt.Scan(&userGuess)
	}	
	fmt.Println("Congratulations! You've guessed the number:", guessingNumber)
}