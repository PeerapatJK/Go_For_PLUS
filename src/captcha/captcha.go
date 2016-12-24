package captcha

func captcha(pattern, firstOperand, secondOperand, operator int) string {

	return operandToText(firstOperand) + "+1"
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
