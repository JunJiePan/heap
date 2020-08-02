package heap

import (
	"sync"
)

type MinHeap struct {
	elements []int
	sync.Mutex
}

func (minHeap *MinHeap) Push(element int) {
	minHeap.Lock()
	defer minHeap.Unlock()
	minHeap.elements = append(minHeap.elements, element)
	index := len(minHeap.elements) - 1
	for {
		parentIndex := (index - 1) / 2
		if parentIndex < 0 || minHeap.elements[parentIndex] <= minHeap.elements[index] {
			break
		}
		minHeap.elements[parentIndex], minHeap.elements[index] = minHeap.elements[index], minHeap.elements[parentIndex]
		index = parentIndex
	}
}

func (minHeap *MinHeap) Pop() int {
	minHeap.Lock()
	defer minHeap.Unlock()
	if len(minHeap.elements) == 0 {
		panic(ErrEmptyHeap)
	}
	delElement := minHeap.elements[0]
	var index int
	for {
		if 2*index+1 > len(minHeap.elements)-1 {
			break
		}
		if 2*index+2 > len(minHeap.elements)-1 {
			minHeap.elements[2*index+1], minHeap.elements[index] = minHeap.elements[index], minHeap.elements[2*index+1]
			index = 2*index + 1
		} else if minHeap.elements[2*index+1] < minHeap.elements[2*index+2] {
			minHeap.elements[2*index+1], minHeap.elements[index] = minHeap.elements[index], minHeap.elements[2*index+1]
			index = 2*index + 1
		} else {
			minHeap.elements[2*index+2], minHeap.elements[index] = minHeap.elements[index], minHeap.elements[2*index+2]
			index = 2*index + 2
		}
	}
	minHeap.elements = append(minHeap.elements[:index], minHeap.elements[index+1:]...)
	return delElement
}
