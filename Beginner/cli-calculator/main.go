package main

import "fmt"

var operators = []string{"+", "-", "*", "/"}

func main() {
	fmt.Println("Welcome to the CLI Calculator!")

	var num1 float64 = 0
	var operator string
	var result *float64
	continueCmd := "y"

	for continueCmd == "y" {
		if result == nil {
			num1 := new(float64)
			fmt.Print("Enter first number: ")
			fmt.Scanln(num1)

			result = num1
		}

		fmt.Print("Enter operator (+, -, *, /): ")
		fmt.Scanln(&operator)

		isValidOperator := validateOperator(operator)

		if ! isValidOperator {
			panic("You used an wrong operator")
		}

		fmt.Print("Enter second number: ")
		fmt.Scanln(&num1)

		calculatedResult := calculate(num1, operator, result)
		result = &calculatedResult
		
		showResult(result) 

		continueCmd = continueOperation()
	} 

	fmt.Println("Exiting the calculator. Goodbye! 😊")
}

func continueOperation() string {
	var continueCmd string
	fmt.Print("Do you want to continue? (y/n or any other touch): ")
	fmt.Scanln(&continueCmd)

	return continueCmd
}

func showResult(result *float64) {
	fmt.Println("=======================================")
	fmt.Println("result: ", *result)
	fmt.Println("========================================")
}

func calculate(num1 float64, operator string, result *float64) float64 {
	if result == nil {
		result = new(float64)
		*result = 0
	}

	switch operator {
		case "+":
			*result += num1
		case "-":
			*result -= num1
		case "*":
			*result *= num1
		case "/":
			if num1 != 0 {
				*result /= num1
			} else {
				fmt.Println("Error: Division by zero is not allowed.")
			}
		default:
			fmt.Println("Error: Invalid operator.")
		}
	return *result
}

func validateOperator(operator string) bool {
	for _, op := range operators {
		if operator == op {
			return true
		}
	}
	return false
}
