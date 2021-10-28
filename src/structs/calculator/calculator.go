package calculator

import (
	"strings"
)

type calculator struct {
	firstNumber  int
	secondNumber int
	operator     string
}

func NewCalculator(firstNumber int, secondNumber int, operator string) *calculator {
	return &calculator{firstNumber, secondNumber, operator}
}

func (calculator calculator) Operation() int {

	switch strings.ToLower(calculator.operator) {

	case "suma":
		return calculator.firstNumber + calculator.secondNumber
	case "resta":
		return calculator.firstNumber - calculator.secondNumber
	case "multiplicacion":
		return calculator.firstNumber * calculator.secondNumber
	case "division":
		return calculator.firstNumber / calculator.secondNumber
	default:
		return -1
	}
}
