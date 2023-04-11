package main

import "fmt"

const x int = 1

const (
	s = "abc"
	y
)

func main() {
	fmt.Println(x, s, y)
}
