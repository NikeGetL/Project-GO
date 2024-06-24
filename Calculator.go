package main

import (
	"fmt"
)

func info1() {
	fmt.Println("Калькулятор с арабскими и римскими числами.\n ")
}

func info2() {
	fmt.Println("Вы можете вводить числа (армабские 1, 2, 3, 4, 5, 6, 7, 8, 9 10), (римские 'I', 'II', 'III', 'IV', 'V', 'VI', 'VII', 'VIII', 'IX', 'X')")
	fmt.Println("и символ вычисления которое вы хотите сделать. Поддерживаются только вычисления сложения(+), вычитания(-), умножения(*) и деления(/).")
	fmt.Println("Вводить нужно в определённой последовательности. Пример: '1 + 1' или \n'1\n +\n 1',\n 'I + I' или \n'I\n +\n I'")
	fmt.Println("И не вводите всё слитно иначе ничего не получится. Вы можете вводить данные соблюдая отступ или перенося данные на новую строку.")
}

func checkOp(op string) string {
	switch op {
	case "+":
	case "-":
	case "*":
	case "/":
	default:
		if op != " " {
			panic("Ошибка была введена строка без необходимых данных и это не является математической операцией")
		} else {
			panic("Ошибка была введенна некорректная операция")
		}

	}
	return op
}

func checkNum(num string) int {
	var num2 int
	switch num {
	case "I":
		num2 = 1
	case "II":
		num2 = 2
	case "III":
		num2 = 3
	case "IV":
		num2 = 4
	case "V":
		num2 = 5
	case "VI":
		num2 = 6
	case "VII":
		num2 = 7
	case "VIII":
		num2 = 8
	case "VIX":
		num2 = 9
	case "X":
		num2 = 10
	case "1":
		num2 = 11
	case "2":
		num2 = 12
	case "3":
		num2 = 13
	case "4":
		num2 = 14
	case "5":
		num2 = 15
	case "6":
		num2 = 16
	case "7":
		num2 = 17
	case "8":
		num2 = 18
	case "9":
		num2 = 19
	case "10":
		num2 = 20
	default:
		if num != " " {
			panic("Ошибка была введена строка без необходимых данных и это не является математической операцией")
		} else {
			panic("Ошибка нет таких чисел для ввода")
		}

	}
	return num2
}

func calculation1(num1 int, op string, num2 int) {
	romanNum := [...]string{"I", "II", "III", "IV", "V", "VI", "VII", "VIII", "IX", "X",
		"XI", "XII", "XIII", "XIV", "XV", "XVI", "XVII", "XVIII", "XIX", "XX",
		"XXI", "XXII", "XXIII", "XXIV", "XXV", "XXVI", "XXVII", "XXVIII", "XXIX", "XXX",
		"XXXI", "XXXII", "XXXIII", "XXXIV", "XXXV", "XXXVI", "XXXVII", "XXXVIII", "XXXIX", "XL",
		"XLI", "XLII", "XLIII", "XLIV", "XLV", "XLVI", "XLVII", "XLVIII", "XLIX", "L",
		"LI", "LII", "LIII", "LIV", "LV", "LVI", "LVII", "LVIII", "LIX", "LX",
		"LXI", "LXII", "LXIII", "LXIV", "LXV", "LXVI", "LXVII", "LXVIII", "LXIX", "LXX",
		"LXXI", "LXXII", "LXXIII", "LXXIV", "LXXV", "LXXVI", "LXXVII", "LXXVIII", "LXXIX", "LXXX",
		"LXXXI", "LXXXII", "LXXXIII", "LXXXIV", "LXXXV", "LXXXVI", "LXXXVII", "LXXXVIII", "LXXXIX", "XC",
		"XCI", "XCII", "XCIII", "XCIV", "XCV", "XCVI", "XCVII", "XCVIII", "XCIX", "C"}
	var result int
	if num1 >= 11 && num2 >= 11 && num1 <= 20 && num2 <= 20 {
		switch op {
		case "+":
			result = (num1 - 10) + (num2 - 10)
			fmt.Println("Результат сложения: ", result)
		case "-":
			result = (num1 - 10) - (num2 - 10)
			fmt.Println("Результат сложения: ", result)
		case "*":
			result = (num1 - 10) * (num2 - 10)
			fmt.Println("Результат сложения: ", result)
		case "/":
			result = (num1 - 10) / (num2 - 10)
			fmt.Println("Результат сложения: ", result)
		}
	} else if num1 >= 1 && num2 >= 1 && num1 <= 10 && num2 <= 10 {
		switch op {
		case "+":
			result = num1 + num2

			fmt.Println("Результат сложения: ", romanNum[result-1])
		case "-":
			result = num1 - num2
			checkCalculat(result)
			fmt.Println("Результат сложения: ", romanNum[result-1])
		case "*":
			result = num1 * num2

			fmt.Println("Результат сложения: ", romanNum[result-1])
		case "/":
			result = num1 / num2
			checkCalculat(result)
			fmt.Println("Результат сложения: ", romanNum[result-1])
		}
	} else {
		panic("Ошибка некорректны ввод данных")
	}
}

func checkCalculat(result int) int {
	if result > 0 {
		return result
	} else {
		panic("Ошибка резултат операции с римскими числами не может быть 0 или быть меньше 0.")
	}
}

func main() {
	var num1, op, num2 string

	info1()
	info2()

	fmt.Scanln(&num1, &op, &num2)

	calculation1(checkNum(num1), checkOp(op), checkNum(num2))
}
