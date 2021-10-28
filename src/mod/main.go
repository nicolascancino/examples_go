package main

import (
	"fmt"

	"github.com/nicolascancino/calculator"
)

func main() {
	c := calculator.NewCalculator(20, 30, "suma")
	response := c.Operation()
	fmt.Println(response)
}
