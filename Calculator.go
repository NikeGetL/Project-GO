package main

import "fmt"

func main() {
	input()
}

func input() {
	var input1, input2, input3 string
	fmt.Println("Input: ")
	fmt.Scanln(&input1, &input2, &input3)

	if input1 == "" && input2 == "" && input3 == "" {
		panic("Empty input.")
	} else if input1 != "" && input2 != "" && input3 == "" {
		if input2[0] == 42 || input2[0] == 43 || input2[0] == 45 || input2[0] == 47 {
			var length int
			length = len(input2)
			if length > 3 {
				panic("Exceeding acceptable values.")
			} else if length < 2 {
				panic("Incomplete data was received.")
			}
			for i := 1; i < length; i++ {
				input3 += string(input2[i])
			}
			if input2[0] == 43 {
				input2 = string(43)
			}
			checkThreeData(input1, input2, input3)
		} else {
			panic("Incomplete data was received.")
		}

	} else if input1 != "" && input2 != "" && input3 != "" {
		var length1, length2, length3 int
		length1 = len(input1)
		length2 = len(input2)
		length3 = len(input3)
		if length1 <= 2 && length2 == 1 && length3 <= 2 {
			checkThreeData(input1, input2, input3)
		} else {
			panic("Exceeding the permissible limit.")
		}

	} else if input1 != "" && input2 == "" && input3 == "" {
		var length int
		length = len(input1)
		if length <= 2 {
			panic("Insufficient data.")
		} else if length >= 3 {
			for i := 0; i < length; i++ {
				if input1[i] == 42 || input1[i] == 43 || input1[i] == 45 || input1[i] == 47 {
					if input1[0] == 48 {
						panic("The data cannot be null")
					}
					if input1[i+1] == 48 {
						panic("The data cannot be null")
					}
				}
			}
			checkOneData(input1)
		}

	}
}

func checkThreeData(num1, op, num2 string) {
	var length1, length2, dataNum1, dataOp, dataNum2, dataAmount1, dataAmount2 int
	dataAmount1 = 48
	dataAmount2 = 41
	length1 = len(num1)
	length2 = len(num2)

	mapRomNum := map[string]int{"I": 1, "II": 2, "III": 3, "IV": 4, "V": 5, "VI": 6, "VII": 7, "VIII": 8, "IX": 9, "X": 10,
		"i": 1, "ii": 2, "iii": 3, "iv": 4, "v": 5, "vi": 6, "vii": 7, "viii": 8, "ix": 9, "x": 10}

	mapArabNum := map[string]int{"1": 1, "2": 2, "3": 3, "4": 4, "5": 5, "6": 6, "7": 7, "8": 8, "9": 9, "10": 10}

	if mapRomNum[num1] > 0 && mapRomNum[num2] > 0 {
		if mapRomNum[num1] <= 10 && mapRomNum[num2] <= 10 {
			dataNum1 = mapRomNum[num1]
			dataNum2 = mapRomNum[num2]
			if op[0] == 42 || op[0] == 43 || op[0] == 45 || op[0] == 47 {

				dataOp = int(op[0]) - dataAmount2
			} else {
				panic("Unavailable actions.")
			}

			calculatorRom(dataNum1, dataOp, dataNum2)
		} else {
			panic("A number greater than 10 was received.")
		}
	} else if mapArabNum[num1] > 0 && mapArabNum[num2] > 0 {
		for i := 0; i < length1; i++ {
			if i == 0 && num1[i] >= 49 && num1[i] <= 57 {
				dataNum1 = int(num1[i]) - dataAmount1
			} else if i == 1 && num1[i] == 48 {
				dataNum1 = 10
			} else {
				panic("There is no correct data in the first value.")
			}
		}
		for i := 0; i < length2; i++ {
			if i == 0 && num1[i] >= 49 && num1[i] <= 57 {
				dataNum2 = int(num2[i]) - dataAmount1
			} else if i == 1 && num2[i] == 48 {
				dataNum2 = 10
			} else {
				panic("There is no correct data in the second value.")
			}
		}

		if op[0] == 42 || op[0] == 43 || op[0] == 45 || op[0] == 47 {

			dataOp = int(op[0]) - dataAmount2
			calculatorArab(dataNum1, dataOp, dataNum2)
		} else {
			panic("Unavailable actions")
		}
	} else {
		panic("Incorrect data was received.")
	}
}

func splitRomStr(data string) {
	var dataNum1, dataNum2, dataOp string
	var a int
	length := len(data)
	a = 0
	for i := 0; i < length; {
		if data[i] == 42 || data[i] == 43 || data[i] == 45 || data[i] == 47 {
			a++
			if a > 1 {
				panic("Exceeding permissible actions.")
			}
			i++
			fmt.Println(a)
		} else {
			i++
		}
	}

	if length <= 9 {

		for i := 0; i < length; i++ {
			if data[i] >= 65 && data[i] <= 90 {
				if dataOp != "" {
					dataNum2 += string(data[i])

				} else {
					dataNum1 += string(data[i])

				}

			} else if data[i] == 42 || data[i] == 43 || data[i] == 45 || data[i] == 47 {
				dataOp += string(data[i])
				if data[i] >= 65 && data[i] <= 90 {
					dataNum2 += string(data[i])
				}
			} else if data[i] >= 97 && data[i] <= 122 {
				if dataOp != "" {
					dataNum2 += string(data[i])
				} else {
					dataNum1 += string(data[i])
				}

			} else if length == 1 {
				panic("Insufficient data.")
			} else {
				panic("Lack of necessary data.")
			}
		}
	} else {
		panic("Exceeding the amount of data received.")
	}
	fmt.Println(dataNum1)
	fmt.Println(dataOp)
	fmt.Println(dataNum2)
	checkThreeData(dataNum1, dataOp, dataNum2)
}

func splitArabStr(data string) {
	var dataNum1, dataNum2, dataOp string
	length := len(data)

	if length <= 5 {
		for i := 0; i < length; i++ {
			if data[i] >= 48 && data[i] <= 71 {
				if dataOp != "" {
					dataNum2 += string(data[i])
				} else {
					dataNum1 += string(data[i])
				}

			} else if data[i] == 42 || data[i] == 43 || data[i] == 45 || data[i] == 47 {
				dataOp += string(data[i])
				if data[i] >= 48 && data[i] <= 71 {
					dataNum2 += string(data[i])
				}
			} else if length == 1 {
				panic("Insufficient data.")
			} else {
				panic("Lack of necessary data.")
			}
		}
	} else {
		panic("Exceeding the amount of data received.")
	}
	checkThreeData(dataNum1, dataOp, dataNum2)
}

func checkOneData(data string) {

	if data[0] >= 65 && data[0] <= 90 || data[0] >= 97 && data[0] <= 122 {
		splitRomStr(data)
	} else {
		splitArabStr(data)
	}
}

func calculatorArab(dataNum1, dataOp, dataNum2 int) {
	var sum int
	if dataOp == 2 {
		sum = dataNum1 + dataNum2
	} else if dataOp == 4 {
		sum = dataNum1 - dataNum2
	} else if dataOp == 1 {
		sum = dataNum1 * dataNum2
	} else if dataOp == 6 {
		sum = dataNum1 / dataNum2
	}
	fmt.Print("Output: ", sum)
}

func calculatorRom(dataNum1, dataOp, dataNum2 int) {
	var sum int
	mapRomNum := map[int]string{
		1: "I", 2: "II", 3: "III", 4: "IV", 5: "V", 6: "VI", 7: "VII", 8: "VIII", 9: "IX", 10: "X",
		11: "XI", 12: "XII", 13: "XIII", 14: "XIV", 15: "XV", 16: "XVI", 17: "XVII", 18: "XVIII", 19: "XIX", 20: "XX",
		21: "XXI", 22: "XXII", 23: "XXIII", 24: "XXIV", 25: "XXV", 26: "XXVI", 27: "XXVII", 28: "XXVIII", 29: "XXIX", 30: "XXX",
		31: "XXXI", 32: "XXXII", 33: "XXXIII", 34: "XXXIV", 35: "XXXV", 36: "XXXVI", 37: "XXXVII", 38: "XXXVIII", 39: "XXXIX", 40: "XL",
		41: "XLI", 42: "XLII", 43: "XLIII", 44: "XLIV", 45: "XLV", 46: "XLVI", 47: "XLVII", 48: "XLVIII", 49: "XLIX", 50: "L",
		51: "LI", 52: "LII", 53: "LIII", 54: "LIV", 55: "LV", 56: "LVI", 57: "LVII", 58: "LVIII", 59: "LIX", 60: "LX",
		61: "LXI", 62: "LXII", 63: "LXIII", 64: "LXIV", 65: "LXV", 66: "LXVI", 67: "LXVII", 68: "LXVIII", 69: "LXIX", 70: "LXX",
		71: "LXXI", 72: "LXXII", 73: "LXXIII", 74: "LXXIV", 75: "LXXV", 76: "LXXVI", 77: "LXXVII", 78: "LXXVIII", 79: "LXXIX", 80: "LXXX",
		81: "LXXXI", 82: "LXXXII", 83: "LXXXIII", 84: "LXXXIV", 85: "LXXXV", 86: "LXXXVI", 87: "LXXXVII", 88: "LXXXVIII", 89: "LXXXIX", 90: "XC",
		91: "XCI", 92: "XCII", 93: "XCIII", 94: "XCIV", 95: "XCV", 96: "XCVI", 97: "XCVII", 98: "XCVIII", 99: "XCIX", 100: "C",
	}
	if dataOp == 2 {
		sum = dataNum1 + dataNum2
	} else if dataOp == 4 {
		sum = dataNum1 - dataNum2
	} else if dataOp == 1 {
		sum = dataNum1 * dataNum2
	} else if dataOp == 6 {
		sum = dataNum1 / dataNum2
	}

	if sum <= 0 {
		panic("The answer in Roman cannot be less than zero.")
	} else {
		fmt.Print("Output: ", mapRomNum[sum])
	}
}
