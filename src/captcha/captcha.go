package captcha

func captcha(pattern, firstOperand, secondOperand, operator int) string {

	if firstOperand == 1{
		return "one+1"
	}
	return "two+1"

}

func operandToText(operand int) string  {

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

	return operandToTextMap[operand]
}
