package main

import "fmt"

func makePrefixZeroes(nums []int) []int {
	prefixZeroes := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		if nums[i-1] == 0 {
			prefixZeroes[i] = prefixZeroes[i-1] + 1
		} else {
			prefixZeroes[i] = prefixZeroes[i-1]
		}
	}
	return prefixZeroes
}

func countZeroes(prefixZeroes []int, l, r int) int {
	return prefixZeroes[r] - prefixZeroes[l]
}

func main() {
	nums := []int{1, 0, 2, 0, 3, 0, 4, 5}
	prefixZeroes := makePrefixZeroes(nums)
	fmt.Println("Префиксные суммы нулей:", prefixZeroes)

	count := countZeroes(prefixZeroes, 1, 7)
	fmt.Println("Количество нулей в диапозоне [1,7):", count)
}
