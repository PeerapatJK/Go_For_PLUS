package main

import "fmt"

const welcome = "Hello"

func main() {
	name := "P"
	surname := "JK"
	fmt.Println(sayHi(welcome, name, surname))
	w, n := sayHi(welcome, name, surname)
	println(w, n)
}
