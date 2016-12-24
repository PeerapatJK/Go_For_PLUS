package main

import "fmt"

func main() {

	fmt.Println("For")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("While")
	var j int;
	for j < 10 {
		fmt.Println(j)
		j++;
	}

	fmt.Println("ForEach")
	ls := []string{"apple","banana","abc"}
	for k,v:= range ls{
		fmt.Println(k,v)
	}

}
