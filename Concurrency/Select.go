package main

import (
	"fmt"
	log "github.com/sirupsen/logrus"
)

func fibonacci(c, quit chan int) {
	x, y := 0, 1
	for {
		select {
		case c <- x:
			log.Info(x)
			x, y = y, x+y
		case <-quit:
			fmt.Println("quit")
			return
		}
	}
}

func main() {
	c := make(chan int)
	quit := make(chan int)
	go func() {
		for i := 0; i < 10; i++ {
			x := <-c
			log.Info(x)
			// fmt.Println(x)
		}
		quit <- 0
	}()
	fibonacci(c, quit)
}
