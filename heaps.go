package main

import "fmt"

// MaxHeap struct has a slice that holds the array
type MaxHeap struct {
	array []int
}

// Insert adds an element to the heap
func (h *MaxHeap) Insert(key int) {
	h.array = append(h.array, key)
	h.maxHeapifyUp(len(h.array) - 1)
}

// Extract returns the largest key, and removes it from the heap
func (h *MaxHeap) Extract() int {
	extracted := h.array[0]
	l := len(h.array) - 1

	if len(h.array) == 0 {
		fmt.Println("cannot extract from empty heap")
		return -1
	}

	h.array[0] = h.array[l]

	h.array = h.array[:l]

	h.maxHeapifyDown(0)

	return extracted
}

// heapify from top to bottom
func (h *MaxHeap) maxHeapifyDown(index int) {
	lastIndex := len(h.array) - 1
	l, r := left(index), right(index)
	childToCompare := 0

	for l <= lastIndex {
		if l == lastIndex {
			childToCompare = l
		} else if h.array[l] > h.array[r] {
			childToCompare = l
		} else {
			childToCompare = r
		}

		if h.array[index] < h.array[childToCompare] {
			h.swap(index, childToCompare)
      index = childToCompare
      l,r = left(index), right(index)
		} else {
      return
    }
	}
}

// heapify from bottom to top
func (h *MaxHeap) maxHeapifyUp(i int) {
	for h.array[parent(i)] < h.array[i] {
		h.swap(parent(i), i)
		i = parent(i)
	}
}

// helpers
// get parent index
func parent(i int) int {
	return (i - 1) / 2
}

// get left child index
func left(i int) int {
	return 2*i + 1
}

// get right child index
func right(i int) int {
	return 2*i + 2
}

// swap keys in the array
func (h *MaxHeap) swap(i, j int) {
	h.array[i], h.array[j] = h.array[j], h.array[i]
}

func main() {
	fmt.Println("Hello World")
	m := &MaxHeap{}
	fmt.Println(m)

	buildHeap := []int{10, 20, 30, 5,7,9,11,13,15,17,500}
	for _, v := range buildHeap {
		m.Insert(v)
		fmt.Println(m)
	}

  for i := 0; i < 5; i++ {
    m.Extract()
    fmt.Println(m)
  }
}
