package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\r\n", s, len(s), cap(s))
}

func main() {
	var s []int
	for i := 0; i < 10; i++ {
		s = append(s, 2*i+1)
		printSlice(s)
	}

	s1 := []int{2, 3, 4, 5}
	fmt.Println("make slice")
	s2 := make([]int, 10, 32)
	printSlice(s2)

	fmt.Println("Copy slice")
	copy(s2, s1)
	printSlice(s2)

	fmt.Println("Delete element from slice")
	s2 = append(s2[:3], s2[4:]...)
	printSlice(s2)
}
