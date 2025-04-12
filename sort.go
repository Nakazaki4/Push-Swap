package main

import (
	"fmt"
	"sort"
)

func (op *StackOperator) Sort() {

	isSorted := false

	for !isSorted {
		op.CyclicSort(op.stackA)

		op.CyclicSort(op.stackB)

		// Check if stack A is sorted
		if op.isStackASorted() {
			fmt.Println(op.stackA.items)
			isSorted = true
		}
	}

}

func (op *StackOperator) CyclicSort(stack *Stack) {
	if stack == op.stackA {
		mid := op.findMid(op.stackA)
		for range op.stackA.Size() {
			top, _ := op.stackA.Peek()
			if top < mid {
				op.PB()
			} else {
				op.RA()
				if op.stackA.Size() == 2 && top > mid {
					op.SA()
					break
				}
			}
		}
	} else if stack == op.stackB {
		mid := op.findMid(op.stackB)
		for range op.stackB.Size() {
			top, _ := op.stackB.Peek()
			if top > mid {
				op.PA()
			} else {
				op.RB()
			}
			if op.stackB.Size() == 2 {
				top, _ = op.stackB.Peek()
				if top > mid {
					op.PA()
					op.PA()
					break
				} else {
					op.SB()
					op.PA()
					op.PA()
				}

			}
		}
	}
}

func (op *StackOperator) CompareTopElements(midA, midB int) {
	topA, _ := op.stackA.Peek()
	topB, _ := op.stackB.Peek()

	if topA
}

func (op *StackOperator) NextElementIndex(stack *Stack, midPoint int) int {
	elements := stack.items
	index := 0
	for i, element := range elements {
		if element < midPoint {
			index = i
			break
		}
	}
	return index
}

func (op *StackOperator) isStackASorted() bool {
	if op.stackA.Size() <= 1 {
		return true
	}

	aCopy := make([]int, op.stackA.Size())
	copy(aCopy, op.stackA.items)

	for i := 0; i < len(aCopy)-1; i++ {
		if aCopy[i] > aCopy[i+1] {
			return false
		}
	}
	return true
}

func (op *StackOperator) findMid(stack *Stack) int {
	s := stack.items
	sort.Ints(s)
	return s[len(s)/2]
}
