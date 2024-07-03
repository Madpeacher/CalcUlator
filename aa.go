package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var rtoar = map[string]int{
	"I": 1, "IV": 4, "V": 5, "IX": 9, "X": 10, "XL": 40, "L": 50, "XC": 90,
	"C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000, // массив, используемый для конвертации римских чисел в арабские в программе калькулятора.
}

var artor = []struct {
	value  int
	symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"}, // для конвертации арабских чисел в римские в программе калькулятора.
}

func rotoint(s string) int {
	n := len(s)          // определяем длину строки s, чтобы знать, сколько символов нужно обработать.
	res := 0             // результат конвертации римских чисел в арабские.
	for i := 0; i < n; { //  цикл, который будет перебирать символы в строке s.
		if i+1 < n { // проверяем, если следующий символ также существует в строке (не вышли за границы).
			if value, true := rtoar[s[i:i+2]]; true { // пытаемся найти двухсимвольное римское число в rtoar. если да, то добавляем его арабское значение (value) к res и увеличиваем i на 2 (чтобы перейти к следующей паре символов).
				res += value
				i += 2
				continue
			}
		}
		//Если не найдено двухсимвольное римское число, переходим к одиночному символу:
		if value, true := rtoar[string(s[i])]; true { // проверяем, есть ли одиночный символ в rtoar. Если есть, добавляем его значение к res и увеличиваем i на 1.
			res += value
			i++
		} else {
			panic("Panic: Invalid Roman numeral") // неверное значение римской цифры
		}
	}
	return res
}

func intor(num int) string { // преобразование арабских чисел в их римское представление, используя массив artor для определения соответствующих значений и символов.
	var res strings.Builder

	for _, entry := range artor { // entry - элемент массива artor
		for num >= entry.value { // value - значение числа в римской системе, symbol - символ рим.сис
			// роверяем, можно ли добавить символ entry.symbol к результату res
			// обрабатывать ввод от пользователя в одной форме и преобразовывать его в другую форму для дальнейшего использования или вывода.
			res.WriteString(entry.symbol)
			num -= entry.value
		}
	}

	return res.String()
}

func calculate(a, b int, op string) int {
	switch op {
	case "+":
		return a + b
	case "-":
		return a - b
	case "*":
		return a * b
	case "/":
		if b == 0 {
			panic("Panic: Division by zero") // делить на 0 нельзя
		}
		return a / b
	default:
		panic("Panic: Invalid operator") // неверный оператор
	}
}

func proinp(input string) string {
	p := strings.Fields(input) // разбивается на отдельные слова
	if len(p) != 3 {           // Проверяется, что строка разбита на ровно три части. Если это условие не выполняется, возвращается ошибка
		panic("Panic: Invalid format") // неверный формат
	}

	a, errA := strconv.Atoi(p[0])
	b, errB := strconv.Atoi(p[2]) // Первый и третий элементы массива  преобразуются из строкового представления чисел в целочисленное
	isR := errA != nil && errB != nil

	if isR {
		a = rotoint(p[0]) // rotoint преобразует римские числа в арабские. Если преобразование невозможно или результат отрицательный, возвращается ошибка.
		b = rotoint(p[2])
	} else {
		if errA != nil || errB != nil {
			panic("Panic: Invalid numbers")
		}

		if (a < 1 || a > 10) || (b < 1 || b > 10) {
			panic("Panic: Numbers must be between 1 and 10")
		}
	}

	res := calculate(a, b, p[1])

	if isR {
		if res <= 0 {
			panic("Panic: Result cannot be negative or zero in Roman numerals")
		}
		return intor(res)
	}

	return strconv.Itoa(res)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите выражение (например, '2 + 3' или 'II * III'): ")
	for scanner.Scan() {
		input := scanner.Text()
		func() {
			defer func() {
				if r := recover(); r != nil {
					fmt.Println(r)
				}
			}()
			res := proinp(input)
			fmt.Println(res)
		}()
		if err := scanner.Err(); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}
}
