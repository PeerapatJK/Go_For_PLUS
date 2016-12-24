package captcha

import (
	"strconv"
	"fmt"
)

func captcha(pattern, firstOperand, secondOperand, operator int) string {

	var a = func() {
		fmt.Println("a")
	}
	a()

	if pattern == 1 {
		return operandToText(firstOperand) + operatorToText(operator) + strconv.Itoa(secondOperand)
	}

	return strconv.Itoa(firstOperand) + operatorToText(operator) + operandToText(secondOperand)
}

type operand struct {
	patternInput, operandInput int
}

func (o *operand) String() string {
	if o.patternInput == 1 {
		return operandToText(o.operandInput)
	}
	return strconv.Itoa(o.operandInput)
}

func operandToText(operand int) string {

	operandToTextMap := map[int]string{
		0:"zero",
		1:"one",
		2:"two",
		3:"three",
		4:"four",
		5:"five",
		6:"six",
		7:"seven",
		8:"eight",
		9:"nine",
	}

	if v, ok := operandToTextMap[operand]; ok {
		return v
	}

	return "Error"
}

func operatorToText(operand int) string {

	operatorToTextMap := map[int]string{
		1:"+",
		2:"-",
		3:"*",
	}

	if v, ok := operatorToTextMap[operand]; ok {
		return v

	}

	return "Error"
}
