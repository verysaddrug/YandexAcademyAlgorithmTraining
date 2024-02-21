package main

import "fmt"

func countZeroSumRanges(nums []int) int {
	count := 0
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j <= len(nums); j++ {
			rangeSum := 0
			for k := i; k < j; k++ {
				rangeSum += nums[k]
			}
			if rangeSum == 0 {
				count++
			}

		}

	}
	return count
}

func main() {
	nums := []int{0, 1, 2, -3, 4, -4}
	fmt.Println("Количество нулевых диапозонов:", countZeroSumRanges(nums))
}
