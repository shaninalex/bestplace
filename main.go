package main

import (
	"fmt"
)

var blocks = []map[string]bool{
	{
		"school": true,  // 0
		"gym":    false, // 1
		"store":  false, // 4
	},
	{
		"school": false, // 1
		"gym":    true,  // 0
		"store":  false, // 3
	},
	{
		"school": true,  // 0
		"gym":    true,  // 0
		"store":  false, // 2
	},
	{ // max distance: 1
		"school": true,  // 0
		"gym":    false, // 1
		"store":  false, // 1
	},
	{
		"school": true,  // 0
		"gym":    false, // 2
		"store":  true,  // 0
	},
}

func BestPlace(blocks []map[string]bool) int {
	var bestPlace int
	for i := range blocks {

		fmt.Printf("Current block: %d\n", i)
		distances := map[string]int{}
		for delta := 1; i-delta >= 0 || i+delta < len(blocks); delta++ {
			if i-delta >= 0 {
				for k, v := range blocks[i-delta] {
					if blocks[i-delta][k] && !blocks[i][k] {
						distances[k] = i - delta
					}
					fmt.Println("Prev: ", k, v, i-delta)
				}
			}
			if i+delta < len(blocks) {
				for k, v := range blocks[i+delta] {
					if blocks[i+delta][k] && !blocks[i][k] {
						distances[k] = i + delta
					}
					fmt.Println("Next: ", k, v, i+delta)
				}
			}
			// TODO: check if all keys in distances > 0 then break
		}
		fmt.Println(distances)
		fmt.Println()
	}
	return bestPlace
}

func main() {
	BestPlace(blocks)
}

