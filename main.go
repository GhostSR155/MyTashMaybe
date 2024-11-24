package main

import (
	"errors"
	"fmt"
	"strconv"
)

func Calc(expression string) (float64, error) {
	stack := make([]float64, 0)
	operators := make([]rune, 0)

	for i, char := range expression {
		if char == ' ' {
			continue
		}
		if char == '(' {
			operators = append(operators, char)
		} else if char == ')' {
			for len(operators) > 0 && operators[len(operators)-1] != '(' {
				stack, operators = applyOperator(stack, operators)
			}
			if len(operators) == 0 || operators[len(operators)-1] != '(' {
				return 0, errors.New("Invalid expression")
			}
			operators = operators[:len(operators)-1]
			stack = stack[:len(stack)-1]
		} else if char == '+' || char == '-' || char == '*' || char == '/' {
			for len(operators) > 0 && precedence(operators[len(operators)-1]) >= precedence(char) {
				stack, operators = applyOperator(stack, operators)
			}
			operators = append(operators, char)
		} else {
			numStr := ""
			for j := i; j < len(expression); j++ {
				if expression[j] >= '0' && expression[j] <= '9' || expression[j] == '.' {
					numStr += string(expression[j])
				} else {
					break
				}
			}
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return 0, errors.New("Invalid number")
			}
			stack = append(stack, num)
		}
	}

	for len(operators) > 0 {
		stack, operators = applyOperator(stack, operators)
	}

	if len(stack) != 1 || len(operators) != 0 {
		return 0, errors.New("Invalid expression")
	}

	return stack[0], nil
}

func precedence(operator rune) int {
	switch operator {
	case '+', '-':
		return 1
	case '*', '/':
		return 2
	}
	return 0
}

func applyOperator(stack []float64, operators []rune) ([]float64, []rune) {
	if len(stack) < 2 || len(operators) == 0 {
		return stack, operators
	}

	b := stack[len(stack)-1]

	a := stack[len(stack)-2]

	stack = stack[:len(stack)-2]

	operator := operators[len(operators)-1]

	operators = operators[:len(operators)-1]

	var result float64
	switch operator {
	case '+':
		result = a + b
	case '-':
		result = a - b
	case '*':
		result = a * b
	case '/':
		if b == 0 {
			return stack, operators
		}
		result = a / b
	}
	stack = append(stack, result)
	return stack, operators
}

func main() {
	result, err := Calc("(2+2*2")
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Result:", result)
	}
}
