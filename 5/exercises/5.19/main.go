package main

import (
	"fmt"
	"math/rand"
)

func main() {
	fmt.Println(noReturn())
}

func noReturn() (result int, err error) {
	defer func() {
		result = recover().(int)
	}()
	panic(rand.Intn(100))
}
