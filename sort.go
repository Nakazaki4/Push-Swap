package main

import "fmt"

func (op *StackOperator) InsertionSort() {
	op.PB()
	op.PB()

	// Sort the first two elements in stack B
	if op.stackB.Size() >= 2 {
		i1, _ := op.stackB.Peek()
		i2 := op.stackB.items[1]
		if i1 < i2 {
			op.SB()
		}

	}

	// Process remaining elements in stack A
	for !op.stackA.IsEmpty() {
		nextVal, _ := op.stackA.Peek()
		position := op.findBestPosition(nextVal)
		op.rotateStackB(position)
		op.PB()
		fmt.Print(op.stackA.items)
		fmt.Print(op.stackB.items)
	}
	// All elements are in the stack B in descending order
	for !op.stackB.IsEmpty() {
		op.PA()
	}
}

func (op *StackOperator) rotateStackB(position int) {
	if position <= op.stackB.Size()/2 {
		// rb
		for i := 0; i < position; i++ {
			op.RB()
		}
	} else {
		// rrb
		for i := 0; i < op.stackB.Size()-position; i++ {
			op.RRB()
		}
	}
}

func (op *StackOperator) findBestPosition(value int) int {
	if op.stackB.IsEmpty() {
		return 0
	}

	// If value is larger than the largest in B, put it at top
	maxVal, maxPos := op.findMaxInStackB()
	if value > maxVal {
		return maxPos
	}

	// If value is smaller than the smallest in B, put it after the smallest
	minVal, minPos := op.findMinInStackB()
	if value < minVal {
		return minPos
	}

	// Otherwise, find the correct position in stack B
	// We want to maintain descending order in B
	for i := 0; i < op.stackB.Size()-1; i++ {
		curr := op.stackB.items[i]
		next := op.stackB.items[i+1]

		// If value fits between curr and next
		if value < curr && value > next {
			return i + 1
		}
	}

	return 0 // Fallback
}

// findMaxInStackB finds the maximum value and its position in stack B
func (op *StackOperator) findMaxInStackB() (int, int) {
	if op.stackB.IsEmpty() {
		return 0, 0
	}

	max := op.stackB.items[0]
	maxPos := 0

	for i, val := range op.stackB.items {
		if val > max {
			max = val
			maxPos = i
		}
	}

	return max, maxPos
}

// findMinInStackB finds the minimum value and its position in stack B
func (op *StackOperator) findMinInStackB() (int, int) {
	if op.stackB.IsEmpty() {
		return 0, 0
	}

	min := op.stackB.items[0]
	minPos := 0

	for i, val := range op.stackB.items {
		if val < min {
			min = val
			minPos = i
		}
	}

	return min, minPos
}
