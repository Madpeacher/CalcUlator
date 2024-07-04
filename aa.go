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
	"C": 100, "CD": 400, "D": 500, "CM": 900, "M": 1000,
}

var artor = []struct {
	value  int
	symbol string
}{
	{1000, "M"}, {900, "CM"}, {500, "D"}, {400, "CD"},
	{100, "C"}, {90, "XC"}, {50, "L"}, {40, "XL"},
	{10, "X"}, {9, "IX"}, {5, "V"}, {4, "IV"}, {1, "I"},
}

func rotoint(s string) int {
	n := len(s)
	res := 0
	for i := 0; i < n; {
		if i+1 < n {
			if value, ok := rtoar[s[i:i+2]]; ok {
				res += value
				i += 2
				continue
			}
		}
		if value, ok := rtoar[string(s[i])]; ok {
			res += value
			i++
		} else {
			panic("Panic: Invalid Roman numeral")
		}
	}
	return res
}

func intor(num int) string {
	var res strings.Builder
	for _, entry := range artor {
		for num >= entry.value {
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
			panic("Panic: Division by zero")
		}
		return a / b
	default:
		panic("Panic: Invalid operator")
	}
}

func proinp(input string) string {
	p := strings.Fields(input)
	if len(p) != 3 {
		panic("Panic: Invalid format")
	}

	a, errA := strconv.Atoi(p[0])
	b, errB := strconv.Atoi(p[2])
	isR := errA != nil && errB != nil

	if isR {
		a = rotoint(p[0])
		b = rotoint(p[2])
	}

	if (a < 1 || a > 10) || (b < 1 || b > 10) {
		panic("Panic: Numbers must be between 1 and 10")
	}

	res := calculate(a, b, p[1])

	if isR {
		romanRes := intor(res)
		if romanRes == "" {
			panic("Panic: Result cannot be negative in Roman numerals")
		}
		return romanRes
	}

	return strconv.Itoa(res)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Введите выражение (например, '2 + 3' или 'II * III'): ")
	for scanner.Scan() {
		input := scanner.Text()
		res := proinp(input)
		fmt.Println(res)
	}
	if err := scanner.Err(); err != nil {
		panic(err)
	}
}
