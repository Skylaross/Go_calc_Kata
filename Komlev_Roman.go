package main

import (
	"fmt"
	"strconv"
)

func main() {
	var aIn, bIn string
	var operator string

	_, err := fmt.Scanf("%s %s %s", &aIn, &operator, &bIn)
	if err != nil {
		fmt.Println("Ошибка: Некорректный формат ввода.\nПример формата X * Y")
		return
	}

	a, b, isRoma, error := ToInt(aIn, bIn)
	if error != "" {
		fmt.Println(error)
		return
	}
	if (a < 1 || a > 10) || (b < 1 || b > 10) {
		fmt.Println("Ошибка: Числа должны быть от 1 до 10")
		return
	}

	var result int
	switch operator {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "/":
		result = a / b
	case "*":
		result = a * b
	default:
		fmt.Println("Ошибка: Допустимые операторы: + - / *")
		return
	}

	if isRoma {
		res := Rims(result)
		fmt.Println(res)
	} else {
		fmt.Println(result)
	}
}

func Rims(x int) string {
	result := ""

	units := []string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX"}
	tens := []string{"X", "XX", "XXX", "XL", "L", "LX", "LXX", "LXXX", "XC"}

	if x == 100 {
		return "C"
	}

	if x/10 > 0 {
		result = result + tens[x/10-1]
		x = x % 10
	}

	if x/1 > 0 {
		result = result + units[x-1]
	}

	return result
}

func ToInt(aIn, bIn string) (int, int, bool, string) {
	a, err := strconv.Atoi(aIn)
	if err != nil {
		a = Rome(aIn)
		if a == 0 {
			return 0, 0, false, "Допускаются только арабские и римские цифры от 1 до 10"
		}

		b := Rome(bIn)
		if a == 0 {
			return 0, 0, false, "Ошибка: Допускаются только арабские и римские цифры от 1 до 10"
		}

		return a, b, true, ""
	}

	b, err := strconv.Atoi(bIn)
	if err != nil {
		return 0, 0, false, "Ошибка: Одновременно можно вводить только  арабские либо римские цифры"
	}

	return a, b, false, ""
}

func Rome(num string) int {
	switch num {
	case "I":
		return 1
	case "II":
		return 2
	case "III":
		return 3
	case "IV":
		return 4
	case "V":
		return 5
	case "VI":
		return 6
	case "VII":
		return 7
	case "VIII":
		return 8
	case "IX":
		return 9
	case "X":
		return 10
	default:
		return 0
	}
}
