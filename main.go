package main

import "fmt"

func main() {
	type firstPearson struct {
		name string
		age  int
	}
	f := firstPearson{name: "John", age: 25}
	var g struct {
		name string
		age  int
	}
	g = f
	fmt.Println(f == g)
}
