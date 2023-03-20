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
		input = strings.TrimSpace(input)
		splittedInput := strings.Split(input, " ")
		// Используем регулярное выражение для проверки на соответствие формату математической операции - два операнда
		//и один оператор (+, -, /, *)
		matched, _ := regexp.MatchString(`\A[0-9VIX]+ [+*\-/] [0-9VIX]+\z`, input)
		//Проверка на количество компонентов строки(Можно было бы это оставить на регулярное выражение но в примерных
		//тестах ошибочные входные данных "1" и "1 + 2 + 3" имеют разные коды ошибок, поэтому использую проверку на длину)
		if len(splittedInput) < 3 {
			//Проверяем не равна ли строка коду выхода из приложения
			if len(splittedInput) == 1 {
				exitCode, _ := strconv.Atoi(input)
				if exitCode == 0 {
					break
				}
			}
			fmt.Println("Вывод ошибки, так как строка не является математической операцией.")
			break
		} else
		// Проверка на соответствие формату математической операции - два операнда и один оператор (+, -, /, *)
		if !matched {
			fmt.Println("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два " +
				"операнда и один оператор (+, -, /, *).")
			break
		}
		// Проверка на римскую систему счисления
		if strings.ContainsAny(input, "VIX") {
			// Оба ли операнда являются числами в римской системе счисления
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

// EvalRoman Обработка математических выражений с числами в римской системе счисления
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

// RomanToArab Перевод чисел из римской СС в арабаскую десятичную
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

// ArabToRoman Перевод чисел из Арабской десятичной СС в Римскую
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
		// Проверка на отрицательность числа
		if strings.Contains(val, "-") {
			// Код ошибки для отрицательного ответа в римской СС: 2
			return res, 2
		} else
		// Проверка на равенство нулю
		if val == "0" {
			// Код ошибки для нулевого ответа в римской СС: 3
			return res, 3
		}
		// Код ошибки для числа вне диапозона [1; 10]
		return res, 1
	}

	return res, 0
}

// EvalArab Обработка математических выражений в Арабской десятичной системе счисления
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
