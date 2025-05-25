package main

import (
	"fmt"
	"sort"
)

type Graph struct {
	possibilities [][]string
}

func (op *StackOperator) Sort() {

	isSorted := false

	for !isSorted {
		op.MidSort(op.stackA)

		op.MidSort(op.stackB)

		// Check if stack A is sorted
		if op.isStackASorted() && op.stackB.IsEmpty() {
			fmt.Println(op.stackA.items)
			isSorted = true
		}
	}
}

func (op *StackOperator) SortThree() {
	items := op.stackA.items
	i1, i2, i3 := items[0], items[1], items[2]

	if i1 > i2 && i2 < i3 && i1 < i3 {
		op.SA()
	} else if i1 > i2 && i2 > i3 {
		op.SA()
		op.RRA()
	} else if i1 > i2 && i2 < i3 && i1 > i3 {
		op.RA()
	} else if i1 < i2 && i2 > i3 && i1 < i3 {
		op.SA()
		op.RA()
	} else if i1 < i2 && i2 > i3 && i1 > i3 {
		op.RRA()
	}
}

func (op *StackOperator) MidSort(stack *Stack) {
	if stack == op.stackA {
		mid := op.findMid(op.stackA)
		for range op.stackA.Size() {
			top, _ := op.stackA.Peek()
			if top < mid {
				op.PB()
			} else {
				if op.stackA.Size() == 3 {
					op.SortThree()
					break
				}
				op.RA()
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
	s := make([]int, len(stack.items))
	copy(s, stack.items)
	sort.Ints(s)
	return s[len(s)/2]
}
