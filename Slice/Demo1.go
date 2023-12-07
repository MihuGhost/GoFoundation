package main

import "fmt"

func main() {
	var s []int

	s1 := make([]int, 12)
	s2 := make([]int, 10, 100)

	s3 := []int{}
	s4 := []int{1, 2, 3}
	printMethod(s, s1, s2, s3, s4)
}

func printMethod(ss ...[]int) {
	for index, s := range ss {
		fmt.Printf("%d  :len: %d  cap: %d\n", index, len(s), cap(s))
	}
}
