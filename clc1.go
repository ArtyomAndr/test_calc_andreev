package main

import (
	"fmt"
)

func main() {

	var num1, num2 float32
	var oper string

	fmt.Print("введи числа")
	fmt.Scan(&num1, &oper, &num2)

	switch oper {

	case "+":
		fmt.Print(num1 + num2)

	case "-":
		fmt.Print(num1 - num2)

	case "*":
		fmt.Print(num1 * num2)

	case "/":
		if num2 == 0 {
			fmt.Print("Не делится на ноль")
		} else {
			fmt.Print(num1 / num2)
		}
	}

}
