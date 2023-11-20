package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	text, _ := reader.ReadString('\n')
	text = strings.Trim(text, "\n")
	arr := strings.Split(text, " ")

	if len(arr) < 3 {
		log.Fatal("Вывод ошибки, так как строка не является математической операцией.")
	} else if len(arr) > 3 {
		log.Fatal("Вывод ошибки, так как формат математической операции не удовлетворяет заданию — два операнда и один оператор (+, -, /, *).")
	} else {
		a := arr[0]
		b := arr[2]
		b = strings.TrimSpace(b)
		znak := arr[1]
		if znak == "+" || znak == "-" || znak == "*" || znak == "/" {
			number_chek_int(a, b, znak)
		} else {
			log.Fatal("Вывод ошибки, так как строка не является математической операцией.")
		}

	}

}

func number_chek_int(a, b, znak string) {

	if x, err := strconv.Atoi(a); err == nil {
		if y, err := strconv.Atoi(b); err == nil {
			if x > 0 && y > 0 && x < 11 && y < 11 {
				result := calculation(x, y, znak)
				fmt.Println(result)
			} else {
				log.Fatal("Число должно быть в диапазоне от 1 до 10 включительно")
			}
		} else {
			log.Fatal("Вывод ошибки, так как используются одновременно разные системы счисления.")
		}

	} else {
		number_chek_rim(a, b, znak)
	}

}

func number_chek_rim(a, b, znak string) {

	numbers := map[string]int{
		"I":    1,
		"II":   2,
		"III":  3,
		"IV":   4,
		"V":    5,
		"VI":   6,
		"VII":  7,
		"VIII": 8,
		"IX":   9,
		"X":    10,
	}

	if _, inMap := numbers[a]; inMap {
		if _, inMap := numbers[b]; inMap {
			result := calculation(numbers[a], numbers[b], znak)
			if result > 0 {
				fmt.Println(convector_rim(result))
			} else {
				log.Fatal("Вывод ошибки, так как в римской системе нет отрицательных чисел.")
			}

		} else {
			log.Fatal("Вывод ошибки, так как используются одновременно разные системы счисления.")
		}

	} else {
		log.Fatal("Вывод ошибки, так как используются одновременно разные системы счисления.")
	}

}

func calculation(a, b int, sign string) int {
	var result int

	switch sign {
	case "+":
		result = a + b
	case "-":
		result = a - b
	case "/":
		result = a / b
	case "*":
		result = a * b
	}
	return result

}

func convector_rim(a int) string {

	var number_rim []string

	for a != 0 {

		if a >= 100 {
			number_rim = append(number_rim, "C")
			a = a - 100
		}
		if a >= 90 && a < 100 {
			number_rim = append(number_rim, "IX")
			a = a - 90
		}
		if a >= 50 && a < 90 {
			number_rim = append(number_rim, "L")
			a = a - 50
		}
		if a >= 40 && a < 50 {
			number_rim = append(number_rim, "XL")
			a = a - 40
		}
		if a >= 10 && a < 40 {
			number_rim = append(number_rim, "X")
			a = a - 10
		}
		if a == 9 {
			number_rim = append(number_rim, "IX")
			a = a - 9
		}
		if a == 4 {
			number_rim = append(number_rim, "IV")
			a = a - 4
		}
		if a < 10 && a >= 5 && a != 9 {
			number_rim = append(number_rim, "V")
			a = a - 5
		}
		if a < 5 && a >= 1 && a != 4 {
			number_rim = append(number_rim, "I")
			a = a - 1
		}

	}
	result := strings.Join(number_rim, "")
	return result
}
