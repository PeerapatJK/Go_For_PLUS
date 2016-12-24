package main

import "fmt"

const welcome = "Hello"

func main() {
	fmt.Println(sayHi())
	w, n := sayHi()
	println(w, n)
}
