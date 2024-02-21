package main

import "fmt"

func countBeatingRooks(rookCoords [][2]int) int {
	addRock := func(rookMap map[int]int, key int) {
		rookMap[key]++
	}

	countPairs := func(rookMap map[int]int) int {
		pairs := 0
		for _, count := range rookMap {
			if count > 1 {
				pairs += count - 1
			}
		}

		return pairs
	}
	rookInRow := make(map[int]int)
	rookInCol := make(map[int]int)

	for _, coord := range rookCoords {
		row, col := coord[0], coord[1]
		addRock(rookInRow, row)
		addRock(rookInCol, col)
	}
	return countPairs(rookInRow) + countPairs(rookInCol)
}

func main() {
	rookCords := [][2]int{{1, 2}, {1, 3}, {2, 2}, {3, 4}}
	fmt.Println("Количество пар ладей, которые могут бить друг друга: ", countBeatingRooks(rookCords))
}
