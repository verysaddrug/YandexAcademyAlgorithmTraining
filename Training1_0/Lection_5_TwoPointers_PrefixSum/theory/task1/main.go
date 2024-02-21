package main

import (
	"fmt"
)

func MakePrefixSum(nums []int) []int {
	prefixSum := make([]int, len(nums)+1)
	for i := 1; i <= len(nums); i++ {
		prefixSum[i] = prefixSum[i-1] + nums[i-1]
	}
	return prefixSum
}

func RSQ(prefixSum []int, l, r int) int {
	if l < 0 || r > len(prefixSum)-1 {
		panic("index out of range")
	}
	return prefixSum[r] - prefixSum[l]
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	prefixSum := MakePrefixSum(nums)
	fmt.Println("Prefix Sums:", prefixSum)

	// Убедитесь, что r не превышает len(nums), поскольку len(prefixSum) = len(nums) + 1
	sum := RSQ(prefixSum, 1, 4) // Сумма элементов с индексами 1 до 3.
	fmt.Println("Range Sum Query [1,4):", sum)
}
