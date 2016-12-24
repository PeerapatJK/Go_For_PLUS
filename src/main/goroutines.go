package main

import (
	"time"
	"fmt"
	"strconv"
	"sync"
)

func say(s string, t int) {
	for i := 1; i <= 5; i++ {

		time.Sleep(time.Duration(1000 * t) * time.Millisecond)
		fmt.Println(s + "  " + strconv.Itoa(i))
	}
}
func main() {
	var wg sync.WaitGroup

	wg.Add(3)
	go say("hello", 1)
	go say("world", 2)
	go func() {
		time.Sleep(2000 * time.Millisecond)
		fmt.Println("Hello World Start ----------------")
	}()

	wg.Wait()
}
