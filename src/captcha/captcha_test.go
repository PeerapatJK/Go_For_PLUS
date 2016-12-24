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

func TestCaptcha_FourMultiply1(t *testing.T) {
	s := captcha(1, 4, 1, 3)
	if s != "four*1" {
		t.Error("we need Four*1 but got : ", s)
	}
}

func TestCaptcha_FourMultiply5(t *testing.T) {
	s := captcha(1, 4, 5, 3)
	if s != "four*5" {
		t.Error("we need Four*5 but got : ", s)
	}
}

func TestCaptcha_6MinusOne(t *testing.T) {
	s := captcha(2, 6, 1, 2)
	if s != "6-one" {
		t.Error("we need 6-One but got : ", s)
	}
}

func TestCaptcha_5MultiplySix(t *testing.T) {
	s := captcha(2, 5, 6, 3)
	if s != "5*six" {
		t.Error("we need 5*six but got : ", s)
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

func TestOperatorToText_WhenGiven1_ShouldReturnPlus(t *testing.T) {
	text := operatorToText(1);
	if text != "+" {
		t.Error("we need + but got : ", text)
	}
}

func TestOperatorToText_WhenGiven3_ShouldReturnMultiply(t *testing.T) {
	text := operatorToText(3);
	if text != "*" {
		t.Error("we need * but got : ", text)
	}
}

