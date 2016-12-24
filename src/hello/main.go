package main

import (
	"fmt"
	"log"
)

const welcome = "Hello"

func main() {
	//name := "P"
	//surname := "JKx"
	w, err := sayHi(welcome)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(w)
}
