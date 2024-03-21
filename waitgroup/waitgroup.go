package main

import (
	"fmt"
	"sync"
)

func main() {

	var wg sync.WaitGroup
	wg.Add(5)

	go func() {
		for i := 1; i <= 5; i++ {
			fmt.Println("正在执行第", i, "次")
			wg.Done()
		}
	}()

	wg.Wait()

	fmt.Println("程序结束")
}
