package main

import (
	"fmt"
)

func findMax2(sl []int) (int, int) {
	max1 := max(sl[0], sl[1])
	max2 := min(sl[0], sl[1])

	for i := 0; i < len(sl); i++ {
		if sl[i] > max1 {
			max2 = max1
			max1 = sl[i]
		} else if sl[i] > max2 {
			max2 = sl[i]
		}
	}
	return max1, max2
}

func main() {
	var n int
	fmt.Println("Введите количество элементов массива:")
	fmt.Scan(&n)
	sl := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < len(sl); i++ {
		fmt.Scan(&sl[i])
	}

	max1, max2 := findMax2(sl)
	fmt.Println(max1, max2)
}
