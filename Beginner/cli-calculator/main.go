package main

import "fmt"

var operators = []string{"+", "-", "*", "/"}

func main() {
	fmt.Println("Welcome to the CLI Calculator!")

	var calculatorStarted bool = false

	var num1 float64 = 0
	var operator string
	var result float64 = 0
	continueCmd := "y"

	for continueCmd == "y" {
		if calculatorStarted == false {
			fmt.Print("Enter first number: ")
			fmt.Scanln(&result)
		} else {
			calculatorStarted = true
		}

		fmt.Print("Enter operator (+, -, *, /): ")
		fmt.Scanln(&operator)

		isValidOperator := validateOperator(operator)

		if ! isValidOperator {
			panic("You used an wrong operator")
		}

		fmt.Print("Enter second number: ")
		fmt.Scanln(&num1)

		result = calculate(num1, operator, result)
		fmt.Printf("Result: %.2f\n", result)

		fmt.Print("Do you want to continue? (y/n or any other touch): ")
		fmt.Scanln(&continueCmd)

		if continueCmd != "y" {
			fmt.Print("Thank for using the calculator. Bye 😊")
		}
	}

}

func calculate(num1 float64, operator string, result float64) float64 {
	switch operator {
		case "+":
			result += num1
		case "-":
			result -= num1
		case "*":
			result *= num1
		case "/":
			if num1 != 0 {
				result /= num1
			} else {
				fmt.Println("Error: Division by zero is not allowed.")
			}
		default:
			fmt.Println("Error: Invalid operator.")
		}
	return result
}

func validateOperator(operator string) bool {
		for _, op := range operators {
			if operator == op {
				return true
			}
		}
		return false
}
