package main

import (
	"fmt"
	"unsafe"
)

func basic() {
	var p *int = nil // 声明*int指针，32位系统4字节，64位系统则是8字节
	var x = 1        // int变量
	p = &x           // &操作符取x的内存地址，赋值给p
	fmt.Println(*p, x)
	*p = 2
	fmt.Println(*p, x)
}

func pointerSize() {
	var p *int = nil // 声明*int指针，32位系统4字节，64位系统则是8字节
	fmt.Println(unsafe.Sizeof(p))
}

type Device struct {
	id   string
	name string
}

func funcPassByPointer() {
	add := func(d Device) {
		fmt.Printf("%p\n", &d)
	}
	d := Device{"1", "d1"}
	fmt.Printf("%p\n", &d)
	add(d)

	addByPointer := func(d *Device) {
		d.name = "d2"
	}
	addByPointer(&d)
	fmt.Println(d)
}

func main() {
	// func comparePointer() {
	num1, num2 := 6, 4
	pt1 := &num1
	pt2 := &num1
	pt3 := &num2

	//只有指向同一个变量，两个指针才相等
	fmt.Printf("%v %v\n", pt1 == pt2, pt1 == pt3) // true false

	fmt.Println("dd", num2)
}
