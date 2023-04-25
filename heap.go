package gdk

import (
	"container/heap"
	"fmt"
)

// heapST to implements the interface of "heap.Interface"
type heapST[E any] struct {
	data []E
	cmp  CMP[E]
}

// implements the methods for "heap.Interface"
func (h *heapST[E]) Len() int { return len(h.data) }
func (h *heapST[E]) Less(i, j int) bool {
	return h.cmp(h.data[i], h.data[j]) < 0
}
func (h *heapST[E]) Swap(i, j int) { h.data[i], h.data[j] = h.data[j], h.data[i] }
func (h *heapST[E]) Push(x any)    { h.data = append(h.data, x.(E)) }
func (h *heapST[E]) Pop() (x any) {
	if len(h.data) <= 0 {
		return x
	}
	// last index
	idx := len(h.data) - 1
	// last item
	x = h.data[idx]
	// update data
	h.data = h.data[:idx]
	// return last item
	return x
}

// Heap base on generics to build a heap tree for any type
type Heap[E any] struct {
	data *heapST[E]
}

// Push push the element x into the heap.
// The complexity is O(log n) where n = h.Len()
func (h *Heap[E]) Push(v E) {
	heap.Push(h.data, v)
}

// Pop remove and return the minimum element(according to Less) from the heap.
// The complexity is O(log n) where n = h.Len()
// Pop is equivalent to Remove(h, 0).
func (h *Heap[E]) Pop() (e E) {
	if h.data.Len() <= 0 {
		return e
	}
	return heap.Pop(h.data).(E)
}

func (h *Heap[E]) Element(index int) (e E, err error) {
	if index < 0 || index >= h.data.Len() {
		return e, fmt.Errorf("out of index")
	}
	return h.data.data[index], nil
}

// Remove remove and return the element at index i from the heap.
// The complexity is O(log n) where n = h.Len()
func (h *Heap[E]) Remove(index int) (e E) {
	if index < 0 || index >= h.data.Len() {
		return e
	}
	return heap.Remove(h.data, index).(E)
}

func (h *Heap[E]) Len() int {
	return len(h.data.data)
}

// Copy to copy heap
func (h *Heap[E]) Copy() *Heap[E] {
	ret := heapST[E]{cmp: h.data.cmp}
	ret.data = make([]E, len(h.data.data))
	copy(ret.data, h.data.data)
	heap.Init(&ret)
	return &Heap[E]{&ret}
}

// NewHeap return Heap pointer and init the heap tree
func NewHeap[E any](t []E, cmp CMP[E]) *Heap[E] {
	ret := heapST[E]{data: t, cmp: cmp}
	heap.Init(&ret)
	return &Heap[E]{&ret}
}
