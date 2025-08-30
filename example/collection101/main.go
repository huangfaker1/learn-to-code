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

    // Demo for collection exercises
    fmt.Println("uniqueStrings:", uniqueStrings([]string{"a", "b", "a", "c", "b"}))
    fmt.Println("intersectInts:", intersectInts([]int{1, 2, 3, 3, 4}, []int{3, 4, 4, 5}))
    fmt.Println("chunkSlice:", chunkSlice([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 3))
    v, ok := safeGet([]string{"x", "y"}, 3)
    fmt.Println("safeGet:", v, ok) // expect: "" false
}

// uniqueStrings deduplicates strings while preserving order.
func uniqueStrings(xs []string) []string {
    seen := make(map[string]struct{}, len(xs))
    out := make([]string, 0, len(xs))
    for _, s := range xs {
        if _, ok := seen[s]; ok {
            continue
        }
        seen[s] = struct{}{}
        out = append(out, s)
    }
    return out
}

// intersectInts returns unique intersection of two int slices.
func intersectInts(a, b []int) []int {
    if len(a) > len(b) {
        a, b = b, a
    }
    setA := make(map[int]struct{}, len(a))
    for _, x := range a {
        setA[x] = struct{}{}
    }
    added := make(map[int]struct{})
    out := make([]int, 0)
    for _, y := range b {
        if _, ok := setA[y]; ok {
            if _, seen := added[y]; !seen {
                out = append(out, y)
                added[y] = struct{}{}
            }
        }
    }
    return out
}

// chunkSlice splits a slice into chunks of given size.
func chunkSlice[T any](xs []T, size int) [][]T {
    if size <= 0 {
        return [][]T{}
    }
    n := len(xs)
    capGuess := n / size
    if n%size != 0 {
        capGuess++
    }
    out := make([][]T, 0, capGuess)
    for i := 0; i < n; i += size {
        end := i + size
        if end > n {
            end = n
        }
        chunk := xs[i:end]
        tmp := make([]T, len(chunk))
        copy(tmp, chunk)
        out = append(out, tmp)
    }
    return out
}

// safeGet returns the i-th element safely.
func safeGet[T any](xs []T, i int) (T, bool) {
    var zero T
    if i < 0 || i >= len(xs) {
        return zero, false
    }
    return xs[i], true
}
