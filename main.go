package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

func calculate(a, b float64, op string) (float64, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("division by zero")
		}
		return a / b, nil
	case "^":
		return math.Pow(a, b), nil
	case "%":
		if b == 0 {
			return 0, fmt.Errorf("modulo by zero")
		}
		return math.Mod(a, b), nil
	default:
		return 0, fmt.Errorf("unknown operator: %s", op)
	}
}

func parse(input string) (float64, string, float64, error) {
	parts := strings.Fields(input)
	if len(parts) != 3 {
		return 0, "", 0, fmt.Errorf("usage: <number> <operator> <number>")
	}
	a, err := strconv.ParseFloat(parts[0], 64)
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid number: %s", parts[0])
	}
	b, err := strconv.ParseFloat(parts[2], 64)
	if err != nil {
		return 0, "", 0, fmt.Errorf("invalid number: %s", parts[2])
	}
	return a, parts[1], b, nil
}

func main() {
	fmt.Println("Simple CLI Calculator")
	fmt.Println("Operators: + - * / ^ %")
	fmt.Println("Type 'exit' to quit")

	scanner := bufio.NewScanner(os.Stdin)
	for {
		fmt.Print("> ")
		if !scanner.Scan() {
			break
		}
		input := strings.TrimSpace(scanner.Text())
		if input == "" {
			continue
		}
		if strings.ToLower(input) == "exit" {
			fmt.Println("Bye!")
			break
		}
		a, op, b, err := parse(input)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		result, err := calculate(a, b, op)
		if err != nil {
			fmt.Println("Error:", err)
			continue
		}
		fmt.Printf("= %g\n", result)
	}
}
