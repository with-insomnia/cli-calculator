package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

func main() {
	first, second, mode, operand, err := arguments() // mode для того чтобы знать какую операцию будем делать римский или арабский
	if err != nil {
		fmt.Print(err)
	}
	math(first, second, mode, operand)
}

// get arguments from terminal
func arguments() (int, int, string, string, error) {
	args := os.Args[1:]
	if len(args) != 3 {
		return 0, 0, "", "", errors.New("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).\nplease try go run . 5 \"*\" 2")
	}
	first, mode1, err := toNum(args[0])
	if err != nil {
		return 0, 0, "", "", err
	}
	second, mode2, err := toNum(args[2])
	if err != nil {
		return 0, 0, "", "", errors.New("Вывод ошибки, так как строка не является математической операцией.")
	}
	oper := args[1]
	if oper != "+" && oper != "-" && oper != "*" && oper != "/" {
		return 0, 0, "", "", errors.New("Вывод ошибки, так как строка не является математической операцией.")
	}
	if mode1 != mode2 {
		return 0, 0, "", "", errors.New("Вывод ошибки, так как используются одновременно разные системы счисления.")
	}
	return first, second, mode1, oper, nil
}

// convert roman to number or  string to arabic
func toNum(str string) (int, string, error) {
	mode := "roman"
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}

	result := 0
	for _, v := range str { // try to convert string to roman numbers
		checkValid := false
		for i, k := range roman {
			if string(v) == k {
				result += numbers[i]
				checkValid = true
			}
		}
		if !checkValid { // if fail try to convet to arabic numbers
			n, err := strconv.Atoi(str)
			if err != nil { // if fail catch  and return error
				return 0, mode, errors.New("Вывод ошибки, так как строка не соответствует требованиям.")
			}
			mode = "arabic"
			return n, mode, nil

		}
	}
	return result, mode, nil
}

// convert number to roman
func toRoman(num int) (string, error) {
	if num == 0 || num >= 4000 {
		return "", errors.New("can't convert")
	}
	numbers := []int{1000, 900, 500, 400, 100, 90, 50, 40, 10, 9, 5, 4, 1}
	roman := []string{"M", "CM", "D", "CD", "C", "XC", "L", "XL", "X", "IX", "V", "IV", "I"}
	ans := ""
	for i := 0; i < len(numbers); i++ {
		if num/numbers[i] != 0 {
			temp := num / numbers[i]
			for k := 1; k <= temp; k++ {
				ans += roman[i]
				num -= numbers[i]
			}
		}
	}
	return ans, nil
}

func math(first int, second int, mode string, operand string) {
	// если у нас римские цифры на входе подается аргументы которые конвертированы в арабский из римского
	// все это для того чтобы выполнить математическу. операцию
	if mode == "roman" {
		_, err := toRoman(first)
		if err != nil {
			fmt.Println(err)
			return
		}
		_, err = toRoman(second)
		if err != nil {
			fmt.Println(err)
			return
		}
		switch operand {
		case "+":
			if first+second > 3999 {
				fmt.Println("Вывод ошибки, так как в римской системе нет цифр больше 3999.")
				return
			}
			res, err := toRoman(first + second)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf(res)
		case "-":
			if first-second < 1 {
				fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
				return
			}
			res, err := toRoman(first - second)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf(res)
		case "*":
			if first*second > 3999 {
				fmt.Println("Вывод ошибки, так как в римской системе нет цифр больше 3999.")
				return
			}
			res, err := toRoman(first * second)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf(res)
		case "/":
			if first/second < 1 {
				fmt.Println("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
				return
			}
			res, err := toRoman(first / second)
			if err != nil {
				fmt.Println(err)
				return
			}
			fmt.Printf(res)
		}
		fmt.Println()
	} else {
		switch operand {
		case "+":
			fmt.Print(first + second)
		case "-":
			fmt.Print(first - second)
		case "*":
			fmt.Print(first * second)
		case "/":
			if second == 0 {
				fmt.Println("Вывод ошибки, так как нельзя делить на ноль.")
				return
			}
			fmt.Print(first / second)
		}
		fmt.Println()
	}
}
