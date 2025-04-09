package main

import "sort"

func (op *StackOperator) Sort() {
	isOneDigit := true

	for _, d := range op.stackA.items {
		if d > 9 {
			isOneDigit = false
		}
	}

	if isOneDigit && op.stackA.Size() <= 6 {
		op.CyclicSort()
	}
}

func (op *StackOperator) CyclicSort() {
	mid := op.findMid()

	for range op.stackA.Size() {
		top, _ := op.stackA.Peek()
		if top < mid {
			op.PB()
		} else {
			op.RA()
		}
	}
	mid = op.findMid()
	

}

func (op *StackOperator) findMid() int {
	s := op.stackA.items
	sort.Ints(s)
	return s[len(s)/2]
}
