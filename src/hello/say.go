package main

func sayHi(sl ...string) (string, string) {
	return sl[0],sl[1]
}
