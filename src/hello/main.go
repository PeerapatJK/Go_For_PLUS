package main

import "fmt"

const welcome = "Hello"

func main() {

	var arr [5]int
	fmt.Println(arr)

	var sl = []int{1,2}
	sl = append(sl,2)
	fmt.Println(sl)

	name := "P"
	age := 1
	fmt.Println(welcome, "Hello", name)
	fmt.Println("age", age)
	//sayHi()
}
