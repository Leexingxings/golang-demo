package main

import "fmt"

func main() {
	u := []int{11, 12, 13, 14, 15}
	fmt.Println("array:", u) // array: [11 12 13 14 15]

	// [12 13]
	s := u[1:3]
	// slice(len=2, cap=4): [12 13]
	fmt.Printf("slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

	s = append(s, 24)
	// after append 24, array: [11 12 13 24 15]
	fmt.Println("after append 24, array:", u)
	// after append 24, slice(len=3, cap=4): [12 13 24]
	fmt.Printf("after append 24, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

	s = append(s, 25)
	// after append 24, array: [11 12 13 24 25]
	fmt.Println("after append 25, array:", u)
	// after append 25, slice(len=4, cap=4): [12 13 24 25]
	fmt.Printf("after append 25, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

	s = append(s, 26)
	// after append 26, array: [11 12 13 24 25]
	fmt.Println("after append 26, array:", u)
	// after append 26, slice(len=5, cap=8): [12 13 24 25 26]
	fmt.Printf("after append 26, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)

	s[0] = 22
	// after reassign 1st elem of slice, array: [11 12 13 24 25]
	fmt.Println("after reassign 1st elem of slice, array:", u)
	// after reassign 1st elem of slice, slice(len=5, cap=8): [22 13 24 25 26]
	fmt.Printf("after reassign 1st elem of slice, slice(len=%d, cap=%d): %v\n", len(s), cap(s), s)
}
