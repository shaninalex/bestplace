package main

import (
	"fmt"
	"slices"
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

	// storing largest distances from features that missing in current block
	var blockDistances []int

	// loop over all blocks
	for i := range blocks {

		// store distances for missing features in currnet block
		distances := map[string]int{}

		// get neighbors from both directions - prev and next
		for delta := 1; i-delta >= 0 || i+delta < len(blocks); delta++ {

			if i-delta >= 0 {
				for k := range blocks[i-delta] {

					// if prev block has feature and current block not
					if blocks[i-delta][k] && !blocks[i][k] {
						_, ok := distances[k]
						if !ok {
							// add delta ( distance )
							distances[k] = delta
						}
					}
				}
			}
			if i+delta < len(blocks) {
				for k := range blocks[i+delta] {

					// if next block has feature and current block not
					if blocks[i+delta][k] && !blocks[i][k] {
						_, ok := distances[k]
						if !ok {
							// add delta ( distance )
							distances[k] = delta
						}
					}
				}
			}
		}
		var dd []int
		for _, d := range distances {
			dd = append(dd, d)
		}
		// append largest distance
		blockDistances = append(blockDistances, slices.Max(dd))
	}

	if len(blockDistances) == 0 {
		return -1
	}

	// return index of block that contain lowest distance of missing features
	minIndex := 0
	minValue := blockDistances[0]

	for i, v := range blockDistances {
		if v < minValue {
			minIndex = i
			minValue = v
		}
	}
	return minIndex
}

func main() {
	fmt.Println(BestPlace(blocks))
}

