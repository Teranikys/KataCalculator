package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	for {
		var result string
		var err int
		fmt.Println("Введите арифметическое выражение, состоящее из двух операндов и одного оператора через пробел.")
		fmt.Println("Введите 0, чтобы выйти из программы:")
		input, _ := reader.ReadString('\n')
		//input = strings.Replace(input, " ", "", -1)
		input = strings.TrimSpace(input)
		splittedInput := strings.Split(input, " ")
		matched, _ := regexp.MatchString(`\A[0-9VIX]+ [+*\-/] [0-9VIX]+\z`, input)
		if len(splittedInput) < 3 {
			if len(splittedInput) == 1 {
				exitCode, _ := strconv.Atoi(input)
				if exitCode == 0 {
					break
				}
			}
			fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
			break
		} else if !matched {
			fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
			break
		}
		if strings.ContainsAny(input, "VIX") {
			if strings.ContainsAny(splittedInput[0], "VIX") != strings.ContainsAny(splittedInput[2], "VIX") {
				fmt.Println("Вывод ошибки, так как используются одновременно разные системы счисления.")
				break
			}
			result, err = EvalRoman(splittedInput)
			if err == 1 {
				fmt.Println("Вывод ошибки, так как числа могут быть только в диапозоне от 1 до 10 включительно, не более")
				break
			} else if err == 2 {
				fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
				break
			} else if err == 3 {
				fmt.Println("Вывод ошибки, так как в римской системе нет нуля.")
				break
			}
		} else {
			result, err = EvalArab(splittedInput)
			if err == 1 {
				fmt.Println("Вывод ошибки, так как числа могут быть только в диапозоне от 1 до 10 включительно, не более")
				break
			}
		}
		fmt.Printf("Результат: %v\n", result)
	}
}

func EvalRoman(exp []string) (string, int) {
	var err int
	exp[0], err = RomanToArab(exp[0])
	if err == 1 {
		return "", err
	}
	exp[2], err = RomanToArab(exp[2])
	if err == 1 {
		return "", err
	}
	num, err := EvalArab(exp)
	if err == 1 {
		return "", err
	}
	res, err := ArabToRoman(num)
	if err != 0 {
		return "", err
	}
	return res, 0
}

func RomanToArab(val string) (string, int) {
	var res string
	switch val {
	case "I":
		res = "1"
	case "II":
		res = "2"
	case "III":
		res = "3"
	case "IV":
		res = "4"
	case "V":
		res = "5"
	case "VI":
		res = "6"
	case "VII":
		res = "7"
	case "VIII":
		res = "8"
	case "IX":
		res = "9"
	case "X":
		res = "10"
	default:
		return res, 1
	}

	return res, 0
}

func ArabToRoman(val string) (string, int) {
	var res string
	switch val {
	case "1":
		res = "I"
	case "2":
		res = "II"
	case "3":
		res = "III"
	case "4":
		res = "IV"
	case "5":
		res = "V"
	case "6":
		res = "VI"
	case "7":
		res = "VII"
	case "8":
		res = "VIII"
	case "9":
		res = "IX"
	case "10":
		res = "X"
	default:
		if strings.Contains(val, "-") {
			return res, 2
		} else if val == "0" {
			return res, 3
		}
		return res, 1
	}

	return res, 0
}

func EvalArab(exp []string) (string, int) {
	x, _ := strconv.Atoi(exp[0])
	y, _ := strconv.Atoi(exp[2])
	if x > 10 || y > 10 {
		return "", 1
	}
	var res string
	switch exp[1] {
	case "+":
		res = strconv.Itoa(x + y)
	case "/":
		res = strconv.Itoa(x / y)
	case "-":
		res = strconv.Itoa(x - y)
	case "*":
		res = strconv.Itoa(x * y)
	}

	return res, 0
}
