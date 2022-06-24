package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	list := []int{1, 2, 3, 4, 5}
	
	for range list {
		rand.Seed(time.Now().UnixNano())
		fmt.Println(rand.Intn(len(list)))
		time.Sleep(time.Duration(30) * time.Millisecond)
	}
}
