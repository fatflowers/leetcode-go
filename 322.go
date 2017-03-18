package main

import "sort"

// assume coins is sorted increasingly
func greedySelect(coins []int, amount int, total int) int {
	for i := len(coins)-1; i>=0; i-- {
		if amount > coins[i] {
			if r := greedySelect(coins[:i+1], amount-coins[i], total+1); r != -1 {
				return r
			}
		} else if amount == coins[i] {
			return total + 1
		}
	}

	return -1
}

func coinChange(coins []int, amount int) int {
	if amount <=0 {
		return 0
	}
	sort.Ints(coins)
	return greedySelect(coins, amount, 0)
}

// func main() {
// 	println(coinChange([]int{1}, 11))
// 	println(coinChange([]int{186,419,83,408}, 6249-419*8 - 408*4 ))
// }