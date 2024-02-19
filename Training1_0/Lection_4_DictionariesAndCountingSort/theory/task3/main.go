package main

import (
	"fmt"
)

func isDigitPermutation(x, y int) bool {
	countDigits := func(num int) [10]int {
		var digitCount [10]int
		for num > 0 {
			lastDigit := num % 10
			digitCount[lastDigit]++
			num /= 10
		}
		return digitCount

	}
	digitsX := countDigits(x)
	digitsY := countDigits(y)

	for i := 0; i < 10; i++ {
		if digitsX[i] != digitsY[i] {
			return false
		}
	}
	return true
}

func main() {
	x := 12345
	y := 54321
	z := 67890

	fmt.Println("x и y являются перестановками цифр друг друга:", isDigitPermutation(x, y))
	fmt.Println("x и z являются перестановками цифр друг друга:", isDigitPermutation(x, z))
}
