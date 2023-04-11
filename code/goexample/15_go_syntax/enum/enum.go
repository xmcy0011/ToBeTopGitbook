package main

import "fmt"

// 基础用法
const (
	x = iota // 0
	y        // 1
)

const (
	_  = iota             // 0
	KB = 1 << (10 * iota) // 1 << (10 * 1)
	MB                    // 1 << (10 * 2), 即MB = 1 << (10 * iota)
)

// 可中断自增值，但是后续自增值包含跳过的行数
const (
	a = iota // 0
	b        // 1, 即b = iota, 0+0=1
	c = 100  // 100
	d        // 100（d省略时，复制上一个值的类型和值，也就是d=100）
	e = iota // 4（显示回复时，自增值包含跳过的2行，所以为1+2+1=4）
	f        // 5
)

type Animal int

// 可自定义类型
const (
	Cat Animal = iota
	Dog
)

type MsgType string

// 也可以是其他类型，比如字符串
const (
	Txt   MsgType = "txt"
	Video MsgType = "video"
	Img           = "img" // 此时不能省略类型，否则变成string类型
)

func main() {
	fmt.Printf("%T %T \n", x, y)
	fmt.Println(KB, MB)
	fmt.Println(a, b, c, d, e, f)
	fmt.Printf("%T: %v - %T: %v \n", Cat, Cat, Dog, Dog)
	fmt.Printf("%T: %v - %T: %v - %T: %v \n", Txt, Txt, Video, Video, Img, Img)
}
