package main

import "fmt"

func main() {
	// chapter3/sources/control_structure_idiom_4.go
	var m = map[string]int{
		"tony": 21,
		"tom":  22,
		"jim":  23,
	}

	counter := 0
	for k, v := range m {
		if counter == 0 {
			//delete(m, "tony")
			m["lucy"] = 24
		}
		counter++
		fmt.Println(k, v)
	}
	fmt.Println("counter is ", counter)

	fmt.Println("last res")
	fmt.Println(m)
}
