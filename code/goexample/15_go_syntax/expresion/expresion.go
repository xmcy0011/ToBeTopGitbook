package main

import (
	"strconv"
)

var x = 0

func ifExample() {
	if x > 0 {
		// ...
	} else if x < 0 {
		// ...
	} else {
		// ...
	}

	if _, err := strconv.Atoi("sd"); err != nil {
		// ...
	}
}

func forExample() {
	// 最常见的循环
	for i := 0; i < 3; i++ {
	}

	// 类似 while x < 10 或 for ; x < 10; x++ {}
	for x < 10 {
		x++
	}

	// 相当于 while true
	for {
		if x > 10 {
			break
		}
	}
}

func forRangeExample() {

}

func main() {

}
