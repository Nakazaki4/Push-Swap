package main

import (
	"sort"
)

var chunks []*Chunk

type Chunk struct {
	Values   []int
	Size     int
	Midpoint int
}

func (op *StackOperator) MidPointSort() {
	if op.stackB.Size() == 0 {
		op.PB()
		op.PB()
	}
	// Sort stack A if it contains only 2 elements
	if op.stackA.Size() == 2 {
		if top, _ := op.stackA.Peek(); top > op.stackA.items[1] {
			op.SA()
		}
		return
	}
	// First we sort the stack A to find the mid value
	midNumber := op.FindMidPoint(op.stackA)
	initialStackASize := op.stackA.Size()
	elementsMoved := 0

	chunk := &Chunk{}
	chunk.Midpoint = midNumber

	for range initialStackASize {
		// If the top number in the stack is greater than midNumber push it to B
		if top, _ := op.stackA.Peek(); top > midNumber {
			op.PB()
			elementsMoved++
			chunk.Values = append(chunk.Values, top)
			// else if the number in the bottom is greater than midNumber  rotate A
		} else if op.stackA.items[op.stackA.Size()-1] > midNumber {
			op.RA()
			// else reverse rotate A
		} else {
			op.RRA()
		}
	}

	if elementsMoved == 0 || elementsMoved == initialStackASize-1 {
		return
	}

	chunk.Size = elementsMoved
	op.AddChunk(chunk)
	op.MidPointSort()

	op.pushFromBtoA()
}

func (op *StackOperator) pushFromBtoA() {
	// To iterate over all the chunks
	for i := len(chunks); i >= 0; i-- {
		chunk := chunks[i]

		// If chunk is already sorted in descending order in stack B,
		// just move everything to A
		if op.IsChunkSorted(chunk) {
			for range chunk.Size {
				op.PA()
			}
			continue
		}

		midNumber := chunk.Midpoint
		// Create a new chunk for elements less than midpoint
		newChunk := &Chunk{}
		newChunk.Midpoint = midNumber

		elementsMoved := 0
		rotationCount := 0

		for range chunk.Size {
			if top, _ := op.stackB.Peek(); top < midNumber {
				op.PA()
				elementsMoved++
				newChunk.Values = append(newChunk.Values, top)
			} else {
				op.RB()
				rotationCount++
			}
		}

		// Restore the rotated elements if needed
		for i:=0; i<rotationCount; i++{
			op.RRB()
		}

		// Add new chunk if elements were moved
        if elementsMoved > 0 {
            newChunk.Size = elementsMoved
            op.AddChunk(newChunk)
            
            // If the new chunk in A has more than 2 elements, sort it
            if elementsMoved > 2 {
                op.SortStackAChunk(newChunk)
            } else if elementsMoved == 2 {
                // Sort 2 elements directly
                if n, _ := op.stackA.Peek(); n > op.stackA.items[1] {
                    op.SA()
                }
            }
        }

		// Handle remaining elements in the original chunk
		remainingSize := chunk.Size - elementsMoved
		if remainingSize > 0 {
			// Create a new chunk for remaining elements
			remainingChunk := &Chunk{
				Size: remainingSize,
				Midpoint: chunk.Midpoint,
			}
			// Process remaining chunk (recursively or directly)
		}
	}
}

func (op *StackOperator) SortStackAChunk(chunk *Chunk){

}

func (op *StackOperator) IsChunkSorted(chunk *Chunk) bool {
	if chunk.Size <= 1 {
		return true
	}

	for i := 0; i < chunk.Size-1; i++ {
		if chunk.Values[i] > chunk.Values[i+1] {
			return false
		}
	}
	return true
}

func (op *StackOperator) AddChunk(chunk *Chunk) {
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
