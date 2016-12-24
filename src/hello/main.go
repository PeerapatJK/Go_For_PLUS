package main

import "fmt"

const welcome = "Hello"

func main() {

	var arr [5]int
	arr[1] = 2
	arr[2] = 5
	arr[3] = 8
	fmt.Println(arr)
	fmt.Println("arr length : ", len(arr))

	var sl = []int{}
	sl = append(sl, 2)
	fmt.Println(sl)
	fmt.Println("sl length : ", len(sl))

	name := "P"
	age := 1
	fmt.Println(welcome, "Hello", name)
	fmt.Println("age", age)
	//sayHi()
}
