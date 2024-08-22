package ex07

import (
	"container/heap"
)

type pair struct {
	n     int
	count int
}

type FreqHeap []pair

func (h FreqHeap) Len() int           { return len(h) }
func (h FreqHeap) Less(i, j int) bool { return h[i].count < h[j].count }
func (h FreqHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *FreqHeap) Push(x any) {
	*h = append(*h, x.(pair))
}

func (h *FreqHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func Solve(ns []int, k int) []int {
	counts := map[int]int{}
	for _, n := range ns {
		counts[n]++
	}

	h := FreqHeap{}
	heap.Init(&h)
	for k, v := range counts {
		if len(h) == k {
			heap.Pop(&h)
		}
		heap.Push(&h, pair{k, v})
	}

	res := []int{}
	for i := 0; i < k; i++ {
		res = append(res, heap.Pop(&h).(pair).n)
	}

	return res
}
