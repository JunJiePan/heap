package heap

import (
	"fmt"
	"testing"
)

func TestMinHeap_Push(t *testing.T) {
	testArr := []int{4, 5, 3, 2, 1, 6}
	var minHeap MinHeap
	for _, v := range testArr {
		minHeap.Push(v)
	}
	fmt.Println(minHeap.elements)
}

func TestMinHeap_Pop(t *testing.T) {
	testArr := []int{4, 5, 3, 2, 1, 6}
	var minHeap MinHeap
	for _, v := range testArr {
		minHeap.Push(v)
	}
	minHeap.Pop()
	fmt.Println(minHeap.elements)
}
