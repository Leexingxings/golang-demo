package main

import (
	"fmt"
	"sync"
)

type Singleton struct {
}

var instance *Singleton
var once sync.Once

func GetInstance() *Singleton {
	once.Do(func() {
		instance = &Singleton{}
	})

	return instance
}

func main() {
	a1 := GetInstance()
	fmt.Printf("%p\n", a1)
	a2 := GetInstance()
	fmt.Printf("%p\n", a2)
}
