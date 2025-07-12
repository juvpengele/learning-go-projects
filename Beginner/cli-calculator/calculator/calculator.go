package calculator

import (
	"fmt"
	"slices"
)

var operators = []string{"+", "-", "*", "/"}

func Calculate(num1 float64, operator string, result *float64) float64 {
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

func ValidateOperator(operator string) bool {
	return slices.Contains(operators, operator)
}