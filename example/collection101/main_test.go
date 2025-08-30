package main

import (
    "reflect"
    "testing"
)

func TestUniqueStrings(t *testing.T) {
    tests := []struct {
        name string
        in   []string
        want []string
    }{
        {"empty", nil, nil},
        {"no-dup", []string{"a", "b", "c"}, []string{"a", "b", "c"}},
        {"has-dup", []string{"a", "b", "a", "c", "b"}, []string{"a", "b", "c"}},
        {"all-dup", []string{"x", "x", "x"}, []string{"x"}},
        {"case-sensitive", []string{"A", "a", "A"}, []string{"A", "a"}},
        {"with-empty", []string{"", "a", "", "b", ""}, []string{"", "a", "b"}},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := uniqueStrings(tc.in)
            if !reflect.DeepEqual(got, tc.want) {
                t.Fatalf("got %v, want %v", got, tc.want)
            }
        })
    }
}

func asSet(xs []int) map[int]struct{} {
    m := make(map[int]struct{}, len(xs))
    for _, v := range xs {
        m[v] = struct{}{}
    }
    return m
}

func TestIntersectInts(t *testing.T) {
    tests := []struct {
        name    string
        a, b    []int
        wantSet map[int]struct{}
    }{
        {"both-empty", nil, nil, asSet(nil)},
        {"one-empty", []int{1, 2}, nil, asSet(nil)},
        {"no-intersection", []int{1, 2}, []int{3, 4}, asSet(nil)},
        {"basic", []int{1, 2, 3, 3, 4}, []int{3, 4, 4, 5}, asSet([]int{3, 4})},
        {"all-dup", []int{1, 1, 1}, []int{1, 1}, asSet([]int{1})},
    }
    for _, tc := range tests {
        t.Run(tc.name, func(t *testing.T) {
            got := intersectInts(tc.a, tc.b)
            if as := asSet(got); !reflect.DeepEqual(as, tc.wantSet) {
                t.Fatalf("got %v (set %v), want set %v", got, as, tc.wantSet)
            }
        })
    }
}

func TestChunkSlice(t *testing.T) {
    t.Run("basic-divisible", func(t *testing.T) {
        got := chunkSlice([]int{1, 2, 3, 4, 5, 6}, 3)
        want := [][]int{{1, 2, 3}, {4, 5, 6}}
        if !reflect.DeepEqual(got, want) {
            t.Fatalf("got %v, want %v", got, want)
        }
    })
    t.Run("basic-not-divisible", func(t *testing.T) {
        got := chunkSlice([]int{1, 2, 3, 4, 5}, 2)
        want := [][]int{{1, 2}, {3, 4}, {5}}
        if !reflect.DeepEqual(got, want) {
            t.Fatalf("got %v, want %v", got, want)
        }
    })
    t.Run("size-le-0", func(t *testing.T) {
        got := chunkSlice([]int{1, 2, 3}, 0)
        if len(got) != 0 {
            t.Fatalf("expect empty, got %v", got)
        }
    })
    t.Run("independent-copies", func(t *testing.T) {
        xs := []int{1, 2, 3, 4}
        chunks := chunkSlice(xs, 2)
        chunks[0][0] = 100
        if xs[0] != 1 {
            t.Fatalf("chunks should not affect original, xs=%v", xs)
        }
    })
    t.Run("generic-string", func(t *testing.T) {
        got := chunkSlice([]string{"a", "b", "c"}, 2)
        want := [][]string{{"a", "b"}, {"c"}}
        if !reflect.DeepEqual(got, want) {
            t.Fatalf("got %v, want %v", got, want)
        }
    })
}

func TestSafeGet(t *testing.T) {
    t.Run("in-bounds", func(t *testing.T) {
        v, ok := safeGet([]int{10, 20, 30}, 1)
        if !ok || v != 20 {
            t.Fatalf("got (%v,%v), want (20,true)", v, ok)
        }
    })
    t.Run("negative-index", func(t *testing.T) {
        _, ok := safeGet([]int{10, 20}, -1)
        if ok {
            t.Fatalf("expect false")
        }
    })
    t.Run("out-of-range", func(t *testing.T) {
        _, ok := safeGet([]int{10, 20}, 5)
        if ok {
            t.Fatalf("expect false")
        }
    })
    t.Run("empty", func(t *testing.T) {
        _, ok := safeGet([]string{}, 0)
        if ok {
            t.Fatalf("expect false")
        }
    })
}

