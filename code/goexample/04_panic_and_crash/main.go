//
// 训练验证：
// - 任意一个 go routine 触发panic，整个程序退出
// - main 也是在一个特殊的g0 routine中执行，同样遵循上述原则
//
package main

import (
	"fmt"
	"log"
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

// 一种技巧，通过 defer func() { recover() }() 来处理异常情况
// 较小的性能消耗：
// 	case1 2189 ms
//  case2 2891 ms
func Example_recover_cost() {
	func1 := func(counter *int) {
		*counter = *counter + 1
	}
	func2 := func(counter *int) {
		defer func() { recover() }()
		*counter = *counter + 1
	}

	t1 := time.Now()
	counter := 0
	for i := 0; i < 1000*1000*1000; i++ {
		func1(&counter)
	}
	log.Println("func1 cost:", time.Now().Sub(t1).Milliseconds(), " ms")

	t1 = time.Now()
	for i := 0; i < 1000*1000*1000; i++ {
		func2(&counter)
	}
	log.Println("func2 cost:", time.Now().Sub(t1).Milliseconds(), " ms")
}
