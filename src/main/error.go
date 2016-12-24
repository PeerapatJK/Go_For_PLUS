package main

import (
	"fmt"
	"encoding/json"
)

func main() {
	errorInverstigate(&myError{EN: "Fail", TH:"เสีย"})
}

func errorInverstigate(err error) {
	fmt.Println(err)

	b,_ := json.Marshal(err)
	fmt.Println(string(b))
}

type myError struct {
	EN string `json:"english"`
	TH string `json:"thai"`
}

func (e *myError) Error() string {
	return "myfail"
}
