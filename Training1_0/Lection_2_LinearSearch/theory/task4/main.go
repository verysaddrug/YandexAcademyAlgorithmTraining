package main

import "fmt"

func findMinEven(sl []int) int {
	res := -1
	for i := 0; i < len(sl); i++ {
		if sl[i]%2 == 0 && (res == -1 || sl[i] < res) {
			res = sl[i]
		}
	}
	return res
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
	res := findMinEven(sl)
	fmt.Println(res)

}
