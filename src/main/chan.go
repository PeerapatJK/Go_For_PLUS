package main

import "fmt"

func say(sch chan string) {
	sch <- "test"
}

func main() {
	mych := make(chan string)
	go say(mych)

	fmt.Println(<-mych)
}
