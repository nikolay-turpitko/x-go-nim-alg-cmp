package main

import (
	"container/heap"
	"runtime"
	"sort"
	"sync"
)

func go_std_sort(a []int) {
	sort.Ints(a)
}

func go_naive_par_sort(a []int) {
	naive_par_sort(intSlice(a))
}

type sortable interface {
	sort.Interface
	Chunk(i, j int) sortable
	LessThenInChunk(i int, chunk sortable, j int) bool
	CopyFromChunk(i int, chunk sortable, j int)
}

type intSlice []int

func (p intSlice) Len() int           { return len(p) }
func (p intSlice) Less(i, j int) bool { return p[i] < p[j] }
func (p intSlice) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p intSlice) Chunk(i, j int) sortable {
	chunk := make([]int, j-i)
	copy(chunk, p[i:j])
	return intSlice(chunk)
}
func (p intSlice) LessThenInChunk(i int, chunk sortable, j int) bool {
	return p[i] < chunk.(intSlice)[j]
}
func (p intSlice) CopyFromChunk(i int, chunk sortable, j int) {
	p[i] = chunk.(intSlice)[j]
}

func naive_par_sort(a sortable) {
	// Split array to chunks and sort them in parallel.
	const (
		minSz     = 100     // threshold to switch to parallel algorithm
		maxSz     = 1000000 // max chunk size
		minCPUNum = 3       // at least 2 CPU for workers
	)

	lna := a.Len()
	numCPU := runtime.NumCPU()

	// For small array (or not enough CPUs) parallelism doesn't worth the trouble.
	if lna < minSz || numCPU < minCPUNum {
		sort.Sort(a)
		return
	}

	// Create workers
	numWorkers := numCPU - 1 // TODO: probably, we don't need to reserve 1 CPU
	toSort := make(chan sortable, numWorkers)
	toMerge := make(chan sortable, numWorkers)
	wg := sync.WaitGroup{}
	wg.Add(numWorkers)
	for i := 0; i < numWorkers; i++ {
		go func() {
			for chunk := range toSort {
				sort.Sort(chunk)
				toMerge <- chunk
			}
			wg.Done()
		}()
	}

	// Wait for workers.
	go func() {
		wg.Wait()
		close(toMerge)
	}()

	// Send job to workers.
	sz := lna / numWorkers // fair share to each worker
	if sz > maxSz {        // but not let them to overwork
		sz = maxSz
	}
	go func() {
		j := sz
		for i := 0; i < lna; i = j {
			chnkSz := sz
			rst := lna - i
			if chnkSz > rst {
				chnkSz = rst
			}
			j = i + chnkSz
			toSort <- a.Chunk(i, j)
		}
		close(toSort)
	}()

	// Merge sorted chunks.
	chunks := make([]sortable, 0, numWorkers)
	posChunk := make([]int, 0, numWorkers)
	for chunk := range toMerge {
		chunks = append(chunks, chunk)
		posChunk = append(posChunk, 0)
	}
	for i := 0; i < lna; i++ {
		fromChnk := -1
		for nchnk, chunk := range chunks {
			if posChunk[nchnk] >= chunk.Len() {
				continue
			}
			if fromChnk == -1 || chunk.LessThenInChunk(posChunk[nchnk], chunks[fromChnk], posChunk[fromChnk]) {
				fromChnk = nchnk
			}
		}
		a.CopyFromChunk(i, chunks[fromChnk], posChunk[fromChnk])
		posChunk[fromChnk]++
	}
}

func go_heap_sort(a []int) {
	h := intHeap(make([]int, len(a)))
	copy(h, a)
	heap.Init(&h)
	a = a[:0]
	for h.Len() > 0 {
		a = append(a, heap.Pop(&h).(int))
	}
}

type intHeap []int

func (h intHeap) Len() int           { return len(h) }
func (h intHeap) Less(i, j int) bool { return h[i] < h[j] }
func (h intHeap) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

func (h *intHeap) Push(x interface{}) {
	// Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.(int))
}

func (h *intHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
