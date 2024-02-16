package main

import "fmt"

func findMax(sl []int) int {
	res := sl[0]
	for i := 0; i < len(sl); i++ {
		if sl[i] > res {
			res = sl[i]
		}
	}
	return res
}

func main() {
	var n int
	fmt.Println("Введите число элементов массива:")
	fmt.Scan(&n)
	sl := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < len(sl); i++ {
		fmt.Scan(&sl[i])
	}

	res := findMax(sl)
	fmt.Printf("Максимальный элемент массива: %d", res)
}
