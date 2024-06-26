package main

import "fmt"

func main() {
	//var input1, input2, input3 string
	const size = 7
	//var quantity1, quantity2, quantity3 int
	//text := make([]string, 0, size)

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
		}
		panic("Incomplete data was received.")
	} else if input1 != "" && input2 != "" && input3 != "" {
		checkThreeData(input1, input2, input3)
	} else if input1 != "" && input2 == "" && input3 == "" {
		var length int
		length = len(input1)
		if length == 2 {
			panic("Not enough yes.")
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
	fmt.Println("Получили данны из трёх: ", num1, op, num2)

}

func checkOneData(data string) {
	fmt.Println("Получили данны из одного: ", data)
}
