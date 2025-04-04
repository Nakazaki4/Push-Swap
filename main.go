package main

import (
	"log"
	"os"
	"strconv"
	"strings"
)

func main(){
	args := os.Args[1]
	Items := strings.Fields(args)
	// Create new stack
	stackA := NewStack("stackA")
	// Push numbers to the stack with the first number at the top of the stack
	for i := len(Items); i > 0; i--{
		n, err := strconv.Atoi(Items[i])
		if err != nil {
			log.Fatal("Error")
		}
		stackA.Push(n)
	}
	// Sorting
}