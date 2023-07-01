package main

import "fmt"

func main() {
	fmt.Println("Hello, 世界")
	fmt.Println("3.0/2.3", 3.0/2.3)
	fmt.Println("true && false", true && false)
	// declaring values
	var x int = 5
	fmt.Println("x is", x)
	// declaring multiple values
	var a, b int = 1, 2
	fmt.Println("a is", a, "b is", b)

	fmt.Println()
	// slice
	s := make([]string, 3)
	fmt.Println("emp:", s)
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("emp:", s)

	fmt.Println()
	fmt.Println("len:", len(s))

	fmt.Println()
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	fmt.Println()
	l := s[2:4]
	fmt.Println("sl1:", l)
	fmt.Println("emp:", s)
}
