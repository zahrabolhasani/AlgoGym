package main

import (
	"fmt"
	"math"
)

func findCheapestPrice(n int, flights [][]int, src int, dst int, k int) int {
	// Initialize the cost array with max values and set the starting point
	cost := make([]int, n)
	for i := range cost {
		cost[i] = math.MaxInt32
	}
	cost[src] = 0

	// Perform at most k + 1 relaxations (allowing k stops means checking k + 1 layers)
	for i := 0; i <= k; i++ {
		tempCost := append([]int{}, cost...) // Copy the current state of costs
		for _, flight := range flights {
			from, to, price := flight[0], flight[1], flight[2]
			// Only update if the starting point has been reached in the previous iteration
			if cost[from] != math.MaxInt32 && cost[from]+price < tempCost[to] {
				tempCost[to] = cost[from] + price
			}
		}
		cost = tempCost
	}

	if cost[dst] == math.MaxInt32 {
		return -1
	}
	return cost[dst]
}

func main() {
	n := 4
	flights := [][]int{
		{0, 1, 100},
		{1, 2, 100},
		{2, 0, 100},
		{1, 3, 600},
		{2, 3, 200},
	}
	src := 0
	dst := 3
	k := 1
	fmt.Println(findCheapestPrice(n, flights, src, dst, k)) // Output: 700
}
