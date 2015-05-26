package main

import "fmt"

func other(in string) string {
	fmt.Println(in)
	return in + " there"
}

func main() {
	a := "hi"
	fmt.Println(a)
	other(a)
}
