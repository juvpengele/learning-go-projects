package main

import (
	"cli-calculator/calculator"
	"fmt"
)

var YES_CMD = "y"

func main() {
	fmt.Println("*********** Welcome to the CLI Calculator! ***********")

	var num2 float64 = 0
	var operator string
	var result *float64
	continueCmd := YES_CMD

	for continueCmd == YES_CMD {
		if result == nil {
			num1 := new(float64)
			fmt.Print("Enter first number: ")
			fmt.Scanln(num1)

			result = num1
		} else{
			calculator.ShowResult("Current Result: ", result)
		}

		fmt.Print("Enter operator (+, -, *, /): ")
		fmt.Scanln(&operator)	

		isValidOperator := calculator.ValidateOperator(operator)

		if ! isValidOperator {
			panic("You used an wrong operator")
		}

		fmt.Print("Enter second number: ")
		fmt.Scanln(&num2)

		calculatedResult := calculator.Calculate(num2, operator, result)
		result = &calculatedResult

		calculator.ShowResult("Result: ", result)

		continueCmd = ContinueOperation()
	}

	fmt.Println("Exiting the calculator. Goodbye! 😊")
}


func ContinueOperation() string {
	var continueCmd string
	fmt.Print("Do you want to continue? (y/n or any other touch): ")
	fmt.Scanln(&continueCmd)

	return continueCmd
}