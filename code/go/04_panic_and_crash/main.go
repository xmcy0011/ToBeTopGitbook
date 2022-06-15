//
// 训练验证：
// - 任意一个 go routine 触发panic，整个程序退出
// - main 也是在一个特殊的g0 routine中执行，同样遵循上述原则
//
package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func Example_main() {
	// func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("panic error:", err)
			fmt.Println("stacktrace from panic:" + string(debug.Stack()))
		}
		fmt.Println("recover")
	}()

	fmt.Println("join party")
	panic("config miss server.ip filed")

	fmt.Println("continues")
	time.Sleep(time.Second * 5)
}

func Example_go() {
	// 捕获 main 所在 go routine
	// defer func() {
	// 	recover()
	// 	fmt.Println("recover")
	// }()

	// 错误：直接使用go关键字创建routine，如果出现异常，整个程序崩溃
	go func() {
		time.Sleep(time.Second * 1)

		// 正确的做法是：为每个routine都执行 recover() 以捕获和恢复
		defer func() {
			recover()
			fmt.Println("recover")
		}()

		// a dengerous action
		panic("throw a error")
	}()

	fmt.Println("server start success")
	time.Sleep(time.Second * 10)
}
