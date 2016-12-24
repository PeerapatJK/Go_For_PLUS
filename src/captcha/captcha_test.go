package captcha

import "testing"

func TestCaptcha_OnePlus1(t *testing.T) {
	s := captcha(1, 1, 1, 1)
	if s != "one+1" {
		t.Error("we need one+1 but got : ", s)
	}
}

func TestCaptcha_TwoPlus1(t *testing.T) {
	s := captcha(1, 2, 1, 1)
	if s != "two+1" {
		t.Error("we need two+1 but got : ", s)
	}
}

func TestOperandToText_WhenGiven2_ShouldReturnTwo(t *testing.T) {
	text := operandToText(2);
	if text != "two" {
		t.Error("we need two but got : ", text)
	}
}

func TestOperandToText_WhenGiven9_ShouldReturnNine(t *testing.T) {
	text := operandToText(9);
	if text != "nine" {
		t.Error("we need nine but got : ", text)
	}
}

