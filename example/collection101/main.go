package main

import (
	"fmt"
	"sort"
)

func Sum(xs []int) int {
	var s int
	for _, v := range xs {
		s += v
	}
	return s
}

func Double(xs []int) []int {
	doubled := make([]int, len(xs))
	for i, v := range xs {
		doubled[i] = v * 2
	}
	return doubled
}

func CharacterCount(s string) (map[rune]int, []rune) {
	counts := make(map[rune]int, len(s))
	for _, r := range s {
		counts[r]++
	}
	keys := make([]rune, 0, len(counts))
	for k := range counts {
		keys = append(keys, k)
	}
	sort.Slice(keys, func(i, j int) bool { return keys[i] < keys[j] })
	return counts, keys
}

// FilterEven returns a new slice containing only the even numbers from the input slice.
// It preallocates the result slice with the same capacity as the input to optimize performance.
func FilterEven(xs []int) []int {
	even := make([]int, 0, len(xs)) // Preallocate capacity to avoid multiple allocations

	for _, v := range xs {
		if v%2 == 0 {
			even = append(even, v)
		}
	}
	return even
}

func main() {
	fmt.Println(FilterEven([]int{1, 2, 3, 4, 5}))
}
