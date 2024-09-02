package main

import (
	"fmt"
	"time"
)

func userWorkFunc(cmd int) error {
	fmt.Println("cmd=", cmd)
	time.Sleep(2 * time.Second)
	return nil
}

func main() {
	gp := NewGoroutinesPool(userWorkFunc, 2)
	for i := 0; i < 10; i++ {
		gp.GoWork(i)
	}
	time.Sleep(10 * time.Second)
}
