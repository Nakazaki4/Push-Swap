package main

import (
	"fmt"
	"sort"
)

func (op *StackOperator) Sort() {

	isSorted := false

	for !isSorted {

		if op.stackA.Size() <= 6 {
			op.CyclicSort(op.stackA)
			fmt.Println(op.stackA.items)
		}
		op.CyclicSort(op.stackB)
		fmt.Println(op.stackB.items)

		// Check if stack A is sorted
		if op.isStackASorted() {
			fmt.Println(op.stackA.items)
			isSorted = true
		}
	}

}

func (op *StackOperator) CyclicSort(stack *Stack) {
	mid := op.findMid()
	if stack == op.stackA {
		for range op.stackA.Size() {
			top, _ := op.stackA.Peek()
			if top < mid {
				op.PB()
			} else {
				op.RA()
			}
			if op.stackA.Size() == 2 && top > mid {
				op.SA()
				break
			}
		}
	} else if stack == op.stackB {
		for range op.stackB.Size() {
			top, _ := op.stackB.Peek()
			if top > mid {
				op.PA()
			} else {
				op.RB()
			}
			if op.stackB.Size() == 2 && top < mid {
				op.SB()
				op.PA()
				op.PA()
			}
		}
	}

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

func (op *StackOperator) findMid() int {
	s := op.stackA.items
	sort.Ints(s)
	return s[len(s)/2]
}
