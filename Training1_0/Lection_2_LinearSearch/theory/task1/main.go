package main

import "fmt"

func findX(sl []int, target int) int {
	res := -1
	for i := 0; i < len(sl); i++ {
		if res == -1 && sl[i] == target {

			res = i
		}
	}
	return res
}

func main() {
	var n, target int
	fmt.Println("Введите количество элементов:")
	fmt.Scan(&n)
	sl := make([]int, n)
	fmt.Println("Введите элементы массива:")
	for i := 0; i < len(sl); i++ {
		fmt.Scan(&sl[i])
	}
	fmt.Println("Введите элемент для поиска")
	fmt.Scan(&target)
	res := findX(sl, target)
	if res != -1 {
		fmt.Printf("Элемент %d найден в массиве под индексом %d \n", target, res)
	} else {
		fmt.Printf("Элемент %d не найден в массиве:", target)
	}

}
