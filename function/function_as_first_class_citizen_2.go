// chapter4/sources/function_as_first_class_citizen_4.go
package main

import "fmt"

func times(x, y int) int {
	return x * y
}

func partialTimes(x int) func(int) int {
	return func(y int) int {
		return times(x, y)
	}
}

func main() {
	timesTwo := partialTimes(2)
	timesThree := partialTimes(3)
	timesFour := partialTimes(4)
	fmt.Println(timesTwo(5))
	fmt.Println(timesThree(5))
	fmt.Println(timesFour(5))
}
