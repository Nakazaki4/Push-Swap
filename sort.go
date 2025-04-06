package main

import (
	"sort"
)

var chunks [][]int

func (op *StackOperator) MidPointSort() {
	if op.stackB.Size() == 0 {
		op.PB()
		op.PB()
	}
	// Sort stack A if it contains only 2 elements
	if op.stackA.Size() == 2 {
		if top, _ := op.stackA.Peek(); top < op.stackA.items[1] {
			op.SA()
		}
		return
	}
	// First we sort the stack A to find the mid value
	midNumber := op.FindMidPoint(op.stackA)
	initialStackASize := op.stackA.Size() - 1
	elementsMoved := 0
	chunk := []int{}
	for i := 0; i < initialStackASize; i++ {
		// If the top number in the stack is less than n push it to B
		if top, _ := op.stackA.Peek(); top < midNumber {
			op.PB()
			elementsMoved++
			chunk = append(chunk, top)
			// else if the number in the bottom is less than n reverse rotate A
		} else if op.stackA.items[op.stackA.Size()-1] < midNumber {
			op.RRA()
			// else just rotate A
		} else {
			op.RA()
		}
	}

	if elementsMoved == 0 || elementsMoved == initialStackASize {
		return
	}

	op.AddChunk(chunk)
	op.MidPointSort()

	
}

func (op *StackOperator) AddChunk(chunk []int) {
	chunks = append(chunks, chunk)
}

func (op *StackOperator) FindMidPoint(stack *Stack) int {
	stackCopy := make([]int, stack.Size()-1)
	copy(stackCopy, op.stackA.items)
	sort.Ints(stackCopy)
	for i, j := 0, len(stackCopy)-1; i < j; i, j = i+1, j-1 {
		stackCopy[i], stackCopy[j] = stackCopy[j], stackCopy[i]
	}
	halfSize := len(stackCopy) / 2
	midNum := stackCopy[halfSize]
	return midNum
}
