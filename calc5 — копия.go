package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	var user string
	var num1, num2 int
	var oper string
	var res int
	var resRom string

	fmt.Print("ввод числа")
	fmt.Scanln(&user)
	under := []rune(user)

	for i := 0; i < len(under); i++ {
		if under[i] == '+' {
			oper = "+"
		}
		if under[i] == '-' {
			oper = "-"
		}
		if under[i] == '/' {
			oper = "/"
		}
		if under[i] == '*' {
			oper = "*"
		}
	}

	underString := string(under)
	block := strings.Split(underString, "[+-/*]")
	stable00 := block[0]
	stable01 := block[1]
	string03 := strings.TrimSpace(stable01)
	num1 = romToNum(stable00)
	num2 = romToNum(string03)
	if num1 < 0 && num2 < 0 {
		res = 0
	} else {
		res = calculated(num1, num2, oper)
		fmt.Println("Римские")
		resRom = convertNumRom(res)
		fmt.Println(stable00 + " " + oper + " " + string03 + "=" + resRom)
	}
	num1, _ = strconv.Atoi(stable00)
	num2, _ = strconv.Atoi(string03)
	res = calculated(num1, num2, oper)
	fmt.Println("Арабские")
	fmt.Println(strconv.Itoa(num1) + " " + oper + " " + strconv.Itoa(num2) + " = " + strconv.Itoa(res))
}

func convertNumRom(numArabian int) string {
	rom := []string{"O", "I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X"}
	return rom[numArabian]
}

func romToNum(rom string) int {
	if rom == "II" {
		return 1
	} else if rom == "II" {
		return 2
	} else if rom == "III" {
		return 3
	} else if rom == "IV" {
		return 4
	} else if rom == "V" {
		return 5
	} else if rom == "VI" {
		return 6
	} else if rom == "VII" {
		return 7
	} else if rom == "VIII" {
		return 8
	} else if rom == "IX" {
		return 9
	} else if rom == "X" {
		return 10
	}
	return -1
}

func calculated(num1, num2 int, op string) int {
	var res int
	switch op {
	case "+":
		res = num1 + num2
	case "-":
		res = num1 - num2
	case "/":
		res = num1 / num2
	case "*":
		res = num1 * num2
	default:
		panic("ошибка")
	}
	return res
}
