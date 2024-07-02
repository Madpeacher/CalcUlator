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

func rotoint(s string) (int, error) {
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
			return 0, fmt.Errorf("Panic") // неверное значение римской цифры
		}
	}
	return res, nil // В конце функции возвращаем res и nil или ошибку в случае неверного ввода римской цифры.
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

func calculate(a, b int, op string) (int, error) {
	switch op {
	case "+":
		return a + b, nil
	case "-":
		return a - b, nil
	case "*":
		return a * b, nil
	case "/":
		if b == 0 {
			return 0, fmt.Errorf("Panic") // делить на 0 нельзя
		}
		return a / b, nil
	default:
		return 0, fmt.Errorf("Panic") // неверный оператор
	}
}

func proinp(input string) (string, error) {
	p := strings.Fields(input) // разбивается на отдельные слова
	if len(p) != 3 {           // Проверяется, что строка разбита на ровно три части. Если это условие не выполняется, возвращается ошибка
		return "", fmt.Errorf("Panic") // неверный формат
	}

	a, errA := strconv.Atoi(p[0])
	b, errB := strconv.Atoi(p[2])     // Первый и третий элементы массива  преобразуются из строкового представления чисел в целочисленное
	isR := errA != nil && errB != nil // Проверяется, являются ли errA и errB не nil. Если оба равны nil, значит, оба числа a и b являются арабскими числами. В противном случае, хотя бы одно из чисел является римским.

	if isR {
		a, errA = rotoint(p[0]) // rotoint преобразует римские числа в арабские. Если преобразование невозможно или результат отрицательный, возвращается ошибка.
		b, errB = rotoint(p[2])
	}

	if errA != nil || errB != nil { // Проверяется, есть ли ошибки в преобразовании чисел. Если хотя бы одно число не удалось преобразовать, возвращается ошибка
		return "", fmt.Errorf("Panic") // неверный формат чисел
	}

	res, err := calculate(a, b, p[1])
	if err != nil {
		return "", err
	}

	if isR {
		romanRes := intor(res)
		if romanRes == "" {
			return "", fmt.Errorf("Panic")
		}
		return romanRes, nil
	}

	return strconv.Itoa(res), nil
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите выражение (например, '2 + 3' или 'II * III'): ")
	for scanner.Scan() {
		input := scanner.Text()
		res, err := proinp(input)
		if err != nil {
			fmt.Println(err)
			return // Прекращаем выполнение при ошибке
		}
		fmt.Println("Результат:", res)
	}
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}
}
