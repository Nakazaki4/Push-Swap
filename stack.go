package main

import "errors"

type Stack struct {
	items []int
	name  string
}

func NewStack(name string) *Stack {
	return &Stack{
		items: []int{},
		name:  name,
	}
}

func (s *Stack) Push(item int) {
	s.items = append([]int{item}, s.items...)
}

func (s *Stack) Pop() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	v := s.items[0]
	s.items = s.items[1:]
	return v, nil
}

func (s *Stack) Peek() (int, error) {
	if len(s.items) == 0 {
		return 0, errors.New("stack is empty")
	}
	return s.items[0], nil
}

func (s *Stack) IsEmpty() bool {
	return len(s.items) == 0
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) IsSorted() bool {
	for i := 0; i < len(s.items)-1; i++ {
		if s.items[i+1] < s.items[i] {
			return false
		}
	}
	return true
}
