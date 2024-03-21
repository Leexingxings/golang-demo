package main

import "fmt"

func main() {
	ch1 := make(chan string)

	ch2 := make(chan struct{})

	go func() {
		ch1 <- "传输数据进来了"
		ch2 <- struct{}{}
	}()

	go func() {
		content, ok := <-ch1
		if !ok {
			fmt.Println("ch1关闭了")
		}
		fmt.Println("接收数据为：", content)
		ch2 <- struct{}{}
	}()

	<-ch2
	<-ch2
	fmt.Println("程序结束")
}
