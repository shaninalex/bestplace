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

func BestPlace(blocks []map[string]bool) {
	for i := range blocks {
		fmt.Printf("Current block: %d\n", i)
		for delta := 1; i-delta >= 0 || i+delta < len(blocks); delta++ {
			if i-delta >= 0 {
				fmt.Printf("Prev: %d (school: %v, gym: %v, store: %v)\n", i-delta, blocks[i-delta]["school"], blocks[i-delta]["gym"], blocks[i-delta]["store"])
			}
			if i+delta < len(blocks) {
				fmt.Printf("Next: %d (school: %v, gym: %v, store: %v)\n", i+delta, blocks[i+delta]["school"], blocks[i+delta]["gym"], blocks[i+delta]["store"])
			}
		}
		fmt.Println()
	}
}

func main() {
	BestPlace(blocks)
}

