package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func rymsToArabicConvert(rymsValue string) int {
	switch rymsValue {
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
	}
	return 0
}

func getRymNumber(number int) string {
	if number < 10 {
		switch number {
		case 1:
			return "I"
		case 2:
			return "II"
		case 3:
			return "III"
		case 4:
			return "IV"
		case 5:
			return "V"
		case 6:
			return "VI"
		case 7:
			return "VII"
		case 8:
			return "VIII"
		case 9:
			return "IX"
		}
		return ""
	}
	switch number {
	case 10:
		return "X"
	case 20:
		return "XX"
	case 30:
		return "XXX"
	case 40:
		return "XL"
	case 50:
		return "L"
	case 60:
		return "LX"
	case 70:
		return "LXX"
	case 80:
		return "LXXX"
	case 90:
		return "XC"
	}
	return ""
}

func arabicToRymsConvert(arabicValue int) string {
	// from 1 to 100
	if arabicValue < 1 {
		return ""
	}
	if arabicValue < 11 {
		return getRymNumber(arabicValue)
	}
	if arabicValue < 100 {
		tens := arabicValue / 10 * 10
		units := arabicValue % 10
		return getRymNumber(tens) + getRymNumber(units)
	}
	return "C"
}

func findSignIndex(inputString string) int {
	index := strings.Index(inputString, "+")
	if index < 1 {
		index = strings.Index(inputString, "-")
	}
	if index < 1 {
		index = strings.Index(inputString, "*")
	}
	if index < 1 {
		index = strings.Index(inputString, "/")
	}
	if len(inputString)-1 == index {
		index = -1
	}
	return index
}

func testOperand(operand string) (int, string) {
	if len(operand) > 4 {
		return 0, ""
	}
	res, err := strconv.Atoi(operand)
	// fmt.Println("Res: ", res, ", err: ", err)
	if err == nil {
		if 0 < res && res < 11 {
			return res, "arabic"
		}
		return 0, ""
	}
	// проверяем на римское
	res2 := rymsToArabicConvert(operand)
	if res2 > 0 {
		return res2, "ryms"
	}
	return 0, ""
}

func calculation(operand1 int, operand2 int, sign byte) int {
	switch sign {
	case 42:
		return operand1 * operand2
	case 43:
		return operand1 + operand2
	case 45:
		return operand1 - operand2
	case 47:
		return operand1 / operand2
	}
	return -1000
}

func main() {
	// ввод данных
	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Введите выражение для вычисления (5+6, 8/3, IX-V):")
	expression, _ := reader.ReadString('\n')

	// обрезаем пробелы
	expression = strings.TrimSpace(expression)
	if expression == "" {
		fmt.Println("Ошибка ввода, введена пустая строка.")
		return
	}

	// Проверяем на наличие знака действия
	signIndex := findSignIndex(expression)
	if signIndex < 1 {
		fmt.Println("Ошибка ввода, отсутствует или не верно расположен знак действия (+, -, *, /).")
		return
	}
	signValue := expression[signIndex]

	//Разделяем по знаку на операнды
	operand1 := strings.TrimSpace(expression[:signIndex])
	operand2 := strings.TrimSpace(expression[signIndex+1:])
	operand1Value, operand1Type := testOperand(operand1) // сделать через структуры
	if operand1Type == "" {
		fmt.Println("Ошибка, неверно задан операнд 1, это должно быть арабское, либо римское число от 1 до 10 включительно.")
		return
	}
	operand2Value, operand2Type := testOperand(operand2)
	if operand2Type == "" {
		fmt.Println("Ошибка, неверно задан операнд 2, это должно быть арабское, либо римское число от 1 до 10 включительно.")
		return
	}
	if operand1Type != operand2Type {
		fmt.Println("Ошибка, несоответствуют типы операндов")
		return
	}

	// вычисление результата и вывод в консоль
	result := calculation(operand1Value, operand2Value, signValue)
	if operand1Type == "arabic" {
		fmt.Println(operand1, string(signValue), operand2, " = ", result)
		return
	}
	rymsResult := arabicToRymsConvert(result)
	if rymsResult == "" {
		fmt.Println("Ошибка, результат вычислений римских чисел меньше единицы.")
		return
	}
	fmt.Println(operand1, string(signValue), operand2, " = ", rymsResult)
}
