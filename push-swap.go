package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {
	osArgs := os.Args
	if len(osArgs) > 2 {
		log.Fatalln("Error")
	}
	args := os.Args[1]
	Items := strings.Fields(args)
	// Create new stack
	stackA := NewStack("stackA")
	// Push numbers to the stack with the first number at the top of the stack
	for i := len(Items)-1; i >= 0; i-- {
		n, err := strconv.Atoi(Items[i])
		if err != nil {
			log.Fatalln("Error")
		}
		stackA.Push(n)
	}
	operator := NewStackOperator(stackA)
	operator.MidPointSort()
	for _, operation := range operator.GetOperations() {
		fmt.Println(operation)
	}
}
