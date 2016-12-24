package main

import "fmt"

const welcome = "Hello"

func main() {
	name := "P"
	fmt.Println(sayHi(name))
	w, n := sayHi(name)
	println(w, n)
}
