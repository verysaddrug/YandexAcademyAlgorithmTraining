package main

import "fmt"

func countPrefixSums(nums []int) map[int]int {
	prefixSumByValue := make(map[int]int)
	nowSum := 0
	prefixSumByValue[nowSum] = 1
	for _, now := range nums {
		nowSum += now
		if _, ok := prefixSumByValue[nowSum]; !ok {
			prefixSumByValue[nowSum] = 0
		}
		prefixSumByValue[nowSum]++
	}
	return prefixSumByValue
}

func countZeroSumZeroes(prefixSumByValue map[int]int) int {
	countRanges := 0
	for _, countSum := range prefixSumByValue {
		countRanges += countSum * (countSum - 1) / 2
	}
	return countRanges
}

func main() {
	nums := []int{0, 1, 2, -3, 4, -4}

	prefixSumByValue := countPrefixSums(nums)

	zeroSumRanges := countZeroSumZeroes(prefixSumByValue)

	fmt.Println("Количество нулевых диапозонов", zeroSumRanges)
}
