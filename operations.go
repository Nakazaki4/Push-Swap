package main

import (
	"errors"
)

type StackOperator struct {
	stackA *Stack
	stackB *Stack
	ops    []string
}

func NewStackOperator(stackA *Stack) *StackOperator {
	return &StackOperator{
		stackA: stackA,
		stackB: NewStack("b"),
		ops:    []string{},
	}
}

func (op *StackOperator) GetOperations() []string {
	return op.ops
}

// push the top first element of stack 'b' to stack 'a'
func (op *StackOperator) PA() error {
	if op.stackB.IsEmpty() {
		return errors.New("stack B is empty")
	}
	num, err := op.stackB.Pop()
	if err != nil {
		return err
	}
	op.stackA.Push(num)
	op.ops = append(op.ops, "pa")
	return nil
}

// Push the top first element of stack 'a' to stack 'b'
func (op *StackOperator) PB() error {
	if op.stackA.IsEmpty() {
		return errors.New("stack A is empty")
	}
	num, err := op.stackA.Pop()
	if err != nil {
		return err
	}
	op.stackB.Push(num)
	op.ops = append(op.ops, "pb")
	return nil
}

// Swap first 2 elements of stack a
func (op *StackOperator) SA() error {
	if op.stackA.Size() < 2 {
		return errors.New("stack A has less than 2 elements")
	}
	op.stackA.items[0], op.stackA.items[1] = op.stackA.items[1], op.stackA.items[0]
	op.ops = append(op.ops, "sa")
	return nil
}

// Swap first 2 elements of stack b
func (op *StackOperator) SB() error {
	if op.stackB.Size() < 2 {
		return errors.New("stack B has less than 2 elements")
	}
	op.stackB.items[0], op.stackB.items[1] = op.stackB.items[1], op.stackB.items[0]
	op.ops = append(op.ops, "sb")
	return nil
}

// Execute 'sa' and 'sb'
func (op *StackOperator) SS() error {
	err1 := op.SA()
	err2 := op.SB()
	if err1 != nil || err2 != nil {
		return errors.New("cannot execute ss")
	}
	op.ops = op.ops[:len(op.ops)-2]
	op.ops = append(op.ops, "ss")
	return nil
}

// Rotates stack 'a'
func (op *StackOperator) RA() error {
	if op.stackA.Size() < 2 {
		return errors.New("stack 'a' has less than 2 elements")
	}
	item, _ := op.stackA.Pop()
	op.stackA.items = append(op.stackA.items, item)
	op.ops = append(op.ops, "ra")
	return nil
}

// Rotates stack 'b'
func (op *StackOperator) RB() error {
	if op.stackB.Size() < 2 {
		return errors.New("stack 'b' has less than 2 elements")
	}
	item, _ := op.stackB.Pop()
	op.stackB.items = append(op.stackB.items, item)
	op.ops = append(op.ops, "rb")
	return nil
}

// Execute both 'ra' and 'rb'
func (op *StackOperator) RR() error {
	err1 := op.RA()
	err2 := op.RB()
	if err1 != nil || err2 != nil {
		return errors.New("cannot execute RR")
	}
	op.ops = op.ops[:len(op.ops)-2]
	op.ops = append(op.ops, "rr")
	return nil
}

// Reverse rotate 'a'
func (op *StackOperator) RRA() error {
	if op.stackA.Size() < 2 {
		return errors.New("stack A has les than 2 elements")
	}
	lastItemIndex := len(op.stackA.items) - 1
	item := op.stackA.items[lastItemIndex]
	op.stackA.items = op.stackA.items[:lastItemIndex]
	op.stackA.Push(item)
	op.ops = append(op.ops, "rra")
	return nil
}

// Reverse rotate 'b'
func (op *StackOperator) RRB() error {
	if op.stackB.Size() < 2 {
		return errors.New("stack B has les than 2 elements")
	}
	lastItemIndex := len(op.stackB.items) - 1
	item := op.stackB.items[lastItemIndex]
	op.stackB.items = op.stackB.items[:lastItemIndex]
	op.stackB.Push(item)
	op.ops = append(op.ops, "rrb")
	return nil
}

func (op *StackOperator) RRR() error {
	err1 := op.RRA()
	err2 := op.RRB()
	if err1 != nil || err2 != nil {
		return errors.New("cannot execute RRR")
	}
	op.ops = op.ops[:len(op.ops)-2]
	op.ops = append(op.ops, "rrr")
	return nil
}
