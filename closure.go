package main

import "fmt"

func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}

func mainClosure() {
	// closure
	nextInt := intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	fmt.Println()
	// new closure
	newInts := intSeq()
	fmt.Println(newInts())
	fmt.Println(newInts())
}
