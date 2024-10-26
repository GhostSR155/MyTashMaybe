package main

import (
	"errors"
	"fmt"
	"strconv"
	"unicode"
)

func Calc(expression string) (float64, error) {
	priority := map[rune]int{'+': 1, '-': 1, '*': 2, '/': 2}
	var num []float64
	var operator []rune
	var hasNumber bool

	for _, ch := range expression {
		if unicode.IsDigit(ch) {
			hasNumber = true
		}
	}

	if !hasNumber {
		return 0, errors.New("no number")
	}

	if len(expression) == 0 {
		return 0, errors.New("empty expression")
	}

	check := rune(expression[len(expression)-1])
	if !unicode.IsDigit(check) && check != ')' {
		return 0, errors.New("invalid, last char is not digits or closing bracket")
	}

	applyOperator := func(a, b float64, op rune) float64 {
		switch op {
		case '+':
			return a + b
		case '-':
			return a - b
		case '*':
			return a * b
		case '/':
			if b == 0 {
				panic("division by zero")
			}
			return a / b
		default:
			return 0
		}
	}

	calculate := func() {
		if len(operator) == 0 || len(num) < 2 {
			return
		}
		b := num[len(num)-1]
		a := num[len(num)-2]
		op := operator[len(operator)-1]
		num = num[:len(num)-2]
		operator = operator[:len(operator)-1]
		result := applyOperator(a, b, op)
		num = append(num, result)
	}

	for i := 0; i < len(expression); i++ {
		ch := rune(expression[i])

		if unicode.IsDigit(ch) || ch == '.' {
			start := i
			for i < len(expression) && (unicode.IsDigit(rune(expression[i])) || expression[i] == '.') {
				i++
			}
			numer, err := strconv.ParseFloat(expression[start:i], 64)
			if err != nil {
				return 0, fmt.Errorf("failed to parse number: %v", err)
			}
			num = append(num, numer)
			i--
		} else if ch == '+' || ch == '-' {
			if i == 0 || expression[i-1] == '(' || len(operator) > 0 && operator[len(operator)-1] == '(' {
				num = append(num, 0)
			}
			for len(operator) > 0 && priority[operator[len(operator)-1]] >= priority[ch] {
				calculate()
			}
			operator = append(operator, ch)
		} else if ch == '*' || ch == '/' {
			for len(operator) > 0 && priority[operator[len(operator)-1]] >= priority[ch] {
				calculate()
			}
			operator = append(operator, ch)
		} else if ch == '(' {
			operator = append(operator, ch)
		} else if ch == ')' {
			for len(operator) > 0 && operator[len(operator)-1] != '(' {
				calculate()
			}
			if len(operator) == 0 {
				return 0, errors.New("mismatched parentheses")
			}
			operator = operator[:len(operator)-1]
		} else if !unicode.IsSpace(ch) {
			return 0, errors.New("invalid character")
		}
	}

	for len(operator) > 0 {
		calculate()
	}

	if len(num) == 1 {
		return num[0], nil
	}

	return 0, errors.New("invalid expression")
}
