package main

import "fmt"

func main() {

	fmt.Println("For---")
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	fmt.Println("While---")
	var j int;
	for j < 10 {
		fmt.Println(j)
		j++;
	}

	fmt.Println("ForEach---")
	ls := []string{"apple", "banana", "abc"}
	for _, v := range ls {
		fmt.Println(v)
	}

	fmt.Println("Map---")
	m := map[string]string{
		"a":"A",
		"b":"B",
		"c":"C",
	}

	for k, v := range m {
		fmt.Println(k, v)
	}

	fmt.Println("SwitchCase---")
	switchCaseTest := 1
	switch switchCaseTest {
	case 0:
		fmt.Println(0)
		fallthrough
	case 1:
		fmt.Println(1)
	default:
		fmt.Println("Not match")
	}

	var a interface{}
	a = "1"

	fmt.Println(a.(string) + "D")

	a = 1
	fmt.Println(a)
}
