package main

import "fmt"

const welcome = "Hello"

func main() {
	name := "P"
	surname := "JK"
	fmt.Println(sayHi(name, surname))
	w, n := sayHi(name, surname)
	println(w, n)
}
