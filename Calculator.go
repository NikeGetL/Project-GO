package main

import (
	"fmt"
)

func info1() {
	fmt.Println("Калькулятор с арабскими и римскими числами.")
	fmt.Println("Выбирете с какими числами вы хотите взаимодействовать.")
	fmt.Println("\nДля выбора введите 1 если хотите работать с арабскими и 2 если хотите работать с римскими числами.")
	fmt.Println("\nЕсть вариант ввода словами, например: 'арабские', 'Арабские', 'арабские числа','Арабские числа',")
	fmt.Println("'римские', 'Римские', 'римские числа', 'Римские числа'.\n ")
}

func info2() {
	fmt.Println("Введите два любых числа от 1 до 10 и символ вычисления которое вы хотите сделать.")
	fmt.Println("Поддерживаются только вычисления сложения(+), вычитания(-), умножения(*) и деления(/).")
	fmt.Println("Вводить нужно в определённой последовательности. Пример: '1 + 1' или \n'1\n +\n 1'")
	fmt.Println("И не вводите всё слитно иначе ничего не получится. Вы можете вводить данные соблюдая отступ или перенося данные на новую строку.")
}

func info3() {
	fmt.Println("")
	fmt.Println("Для ввода доступны римские числа 'I, II, III, IV, V, VI, VII, VIII, VIX и X' и символ вычисления которое вы хотите сделать.")
	fmt.Println("Поддерживаются только вычисления сложения(+), вычитания(-), умножения(*) и деления(/).")
	fmt.Println("Вводить нужно в определённой последовательности. Пример: 'I + I' или \n'I\n +\n I'")
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

func checkNum1(num int) int {
	switch num {
	case 1:
	case 2:
	case 3:
	case 4:
	case 5:
	case 6:
	case 7:
	case 8:
	case 9:
	case 10:
	default:
		if num != ' ' {
			panic("Ошибка была введена строка без необходимых данных и это не является математической операцией")
		} else {
			panic("Ошибка было введенно некорректное число")
		}

	}
	return num
}

func checkNum2(num string) int {
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
	default:
		if num != " " {
			panic("Ошибка была введена строка без необходимых данных и это не является математической операцией")
		} else {
			panic("Ошибка нет таких чисел для ввода")
		}

	}
	return num2
}

func calculation1(num1 int, op string, num2 int) int {
	var result int
	switch op {
	case "+":
		result = num1 + num2
	case "-":
		result = num1 - num2
	case "*":
		result = num1 * num2
	case "/":
		result = num1 / num2
	}
	return result
}

func calculation2(num1 int, op string, num2 int) {
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

}

func checkCalculat(result int) int {
	if result > 0 {
		return result
	} else {
		panic("Ошибка резултат операции с римскими числами не может быть 0 или быть меньше 0.")
	}
}

func main() {
	var choice string
	info1()

	fmt.Scanln(&choice)

	if choice == "1" || choice == "арабские" || choice == "Арабские" || choice == "арабские числа" || choice == "Арабские числа" {
		fmt.Println("Вы выбрали арабские.\n ")
		info2()
		var op string
		var num1, num2 int
		fmt.Scanln(&num1, &op, &num2)
		fmt.Println("Результат: ", calculation1(checkNum1(num1), checkOp(op), checkNum1(num2)))

	} else if choice == "2" || choice == "римские" || choice == "Римские" || choice == "римские числа" || choice == "Римские числа" {
		fmt.Println("Вы выбрали римские.\n ")
		var num1, op, num2 string
		info3()
		fmt.Scanln(&num1, &op, &num2)
		calculation2(checkNum2(num1), checkOp(op), checkNum2(num2))

	} else if choice != " " {
		panic("Ошибка была введёна строка без необходимых данных")

	} else {
		panic("Ошибка было введено неверное значение")
	}
}
