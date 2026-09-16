package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	firstStr := scanner.Text()
	first, err := strconv.Atoi(firstStr)
	if err != nil {
		fmt.Println("Invalid first operand")
		return
	}
	scanner.Scan()
	secondStr := scanner.Text()
	second, err := strconv.Atoi(secondStr)
	if err != nil {
		fmt.Println("Invalid second operand")
		return
	}
	scanner.Scan()
	op := strings.TrimSpace(scanner.Text())
	switch op {
	case "+":
		fmt.Println(first + second)
	case "-":
		fmt.Println(first - second)
	case "*":
		fmt.Println(first * second)
	case "/":
		if second == 0 {
			fmt.Println("Division by zero")
			return
		}
		fmt.Println(first / second)
	default:
		fmt.Println("Invalid operation")
	}
}
