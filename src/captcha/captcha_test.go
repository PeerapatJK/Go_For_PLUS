package captcha

import "testing"

func TestOnePlus1(t *testing.T) {
	s := captcha(1, 1, 1, 1)
	if s != "one+1" {
		t.Error("we need one+1 but got : ", s)
	}
}

func TestTwoPlus1(t *testing.T) {
	s := captcha(1, 2, 1, 1)
	if s != "two+1" {
		t.Error("we need two+1 but got : ", s)
	}
}
