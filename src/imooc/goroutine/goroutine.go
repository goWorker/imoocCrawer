package main

import (
	"fmt"
	"time"
)

func main() {

	for i := 0; i < 1000; i++ {
		go func(i int) {
			for {
				//a[i]++
				//runtime.Gosched()   //手动交出控制权
				fmt.Printf("Hello from goroutine %d\n", i)
			}
		}(i)
	}
	time.Sleep(time.Minute)

}
