package main

import (
	"fmt"
	"sort"
)

// Item has a value and a weight; fractions of it may be taken.
type Item struct {
	Value  int
	Weight int
}

// MaxValue fills a knapsack of the given capacity, taking items in
// descending value-per-weight order and a final fraction of the item that
// no longer fits whole.
func MaxValue(items []Item, capacity int) float64 {
	sorted := append([]Item{}, items...)
	sort.Slice(sorted, func(i, j int) bool {
		// Compare value/weight without division: cross-multiply.
		return sorted[i].Value*sorted[j].Weight > sorted[j].Value*sorted[i].Weight
	})
	total := 0.0
	remaining := capacity
	for _, it := range sorted {
		if remaining == 0 {
			break
		}
		if it.Weight <= remaining {
			total += float64(it.Value)
			remaining -= it.Weight
		} else {
			total += float64(it.Value) * float64(remaining) / float64(it.Weight)
			remaining = 0
		}
	}
	return total
}

func main() {
	fmt.Println(MaxValue([]Item{{60, 10}, {100, 20}, {120, 30}}, 50))
}
