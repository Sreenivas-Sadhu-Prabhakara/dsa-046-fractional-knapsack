package main

import (
	"math"
	"testing"
)

func TestMaxValue(t *testing.T) {
	cases := []struct {
		items    []Item
		capacity int
		want     float64
	}{
		{[]Item{{60, 10}, {100, 20}, {120, 30}}, 50, 240},
		{[]Item{{500, 30}}, 10, 500.0 / 3},
		{[]Item{{10, 5}}, 0, 0},
		{[]Item{}, 10, 0},
		{[]Item{{10, 2}, {10, 2}}, 10, 20}, // capacity to spare
	}
	for _, c := range cases {
		if got := MaxValue(c.items, c.capacity); math.Abs(got-c.want) > 1e-9 {
			t.Errorf("MaxValue(%v,%d) = %v, want %v", c.items, c.capacity, got, c.want)
		}
	}
}
