package main

import (
	"testing"
)

func TestCalculator(t *testing.T) {
	//Test 1
	input := "1 + 2"
	expected := "3"

	result, _ := StringProccessing(input)

	if result != expected {
		t.Errorf("Неверный результат. Ожидалось %v, получилось %v",
			expected, result)
	}

	//Test 2
	input = "VI / III"
	expected = "II"

	result, _ = StringProccessing(input)

	if result != expected {
		t.Errorf("Неверный результат. Ожидалось %v, получилось %v",
			expected, result)
	}

	//Test 3
	input = "I - II"
	expected = "Вывод ошибки, так как в римской системе нет отрицательных чисел."

	result, _ = StringProccessing(input)

	if result != expected {
		t.Errorf("Неверный результат. Ожидалось \n%v, получилось \n%v",
			expected, result)
	}

	//Test 4
	input = "I + 1"
	expected = "Вывод ошибки, так как используются одновременно разные системы счисления."

	result, _ = StringProccessing(input)

	if result != expected {
		t.Errorf("Неверный результат. Ожидалось \n%v, получилось \n%v",
			expected, result)
	}

	//Test 5
	input = "1"
	expected = "Вывод ошибки, так как строка не является математической операцией."

	result, _ = StringProccessing(input)

	if result != expected {
		t.Errorf("Неверный результат. Ожидалось \n%v, получилось \n%v",
			expected, result)
	}

	//Test 6
	input = "1 + 2 + 3"
	expected = "Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *)."

	result, _ = StringProccessing(input)

	if result != expected {
		t.Errorf("Неверный результат. Ожидалось \n%v, получилось \n%v",
			expected, result)
	}
}
