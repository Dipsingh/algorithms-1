package symboltables

import (
	"fmt"
)

type BSTElem struct {
	key   int32
	value int32
}

type BinarySearchST struct {
	array []BSTElem
}

func (bss *BinarySearchST) Append(key, value int32) {
	bss.array = append(bss.array, BSTElem{key: key, value: value})
}

func (bss *BinarySearchST) Print() {
	for cntr := 0; cntr < len(bss.array); cntr++ {
		fmt.Println("key: ", bss.array[cntr].key, " value: ", bss.array[cntr].value)
	}
}

func (bss *BinarySearchST) Len() int {
	return len(bss.array)
}

func (bss *BinarySearchST) Less(i, j int) bool {
	return bss.array[i].key < bss.array[j].key
}

func (bss *BinarySearchST) Swap(i, j int) {
	bss.array[i], bss.array[j] = bss.array[j], bss.array[i]
}

func (bss *BinarySearchST) Search(key int32) int32 {
	low := 0
	high := bss.Len() - 1
	for low <= high {
		mid := low + (high-low)/2
		if bss.array[mid].key > key {
			high = mid - 1
		} else if bss.array[mid].key < key {
			low = mid + 1
		} else {
			return bss.array[mid].value
		}
	}
	return -1
}
