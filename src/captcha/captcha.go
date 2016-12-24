package captcha

import (
	"strconv"
	"log"
	"fmt"
)

func captcha(pattern, firstOperand, secondOperand, operator int) string {

	firstConvertedOperand := operand{
		isText: pattern == 1,
		operandInput: firstOperand,
	}

	secondConvertedOperand := operand{
		isText: pattern == 2,
		operandInput: secondOperand,
	}

	defer func() {
		fmt.Println("End.")
	}()

	return firstConvertedOperand.String() + operatorToText(operator) + secondConvertedOperand.String()
}

var Captcha = captcha

type operand struct {
	isText bool
	operandInput int
}

func (o *operand) String() string {
	if o.isText == true {
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

	log.Fatal("Value not found in operandToTextMap")

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

	log.Fatal("Value not found in operatorToTextMap")

	return "Error"
}
