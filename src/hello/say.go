package main

import "errors"

func sayHi(sl ...string) (string, error) {
	if len(sl) < 3 {
		return "", errors.New("We need at least 3 parameters")
	}
	return sl[0] + sl[1] + sl[2], nil
}
