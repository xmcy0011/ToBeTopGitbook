package main

import "fmt"

func ParamLifeFunc(p *int) {
	go func() {
		println(p)
	}()
}

// TestParamLifeFunc 延长P生命周期，经过内存逃逸分析，最终p被分配到了堆上（原本应该分配在栈上）
// PS：函数名先改成main再执行下面的命令
// 1) 使用 go build -gcflags="-m -l" 编译，-m 输出编译器优化策略，-l 禁止函数内联，可以看到 moved to heap: x，说明x逃逸了，最终在堆上分配了内存
// 2) 使用 go tool objdump -s "main.main" 15_go_syntax 查看反汇编代码，看到 runtime.newobject(SB) 说明在堆上为x分配了内存
func TestParamLifeFunc() {
	x := 100
	p := &x
	ParamLifeFunc(p)
}

func HideClosureFunc(x int) func() {
	return func() {
		fmt.Println(x)
	}
}

// func TestHideClosureFunc(){
func main() {
	f := HideClosureFunc(100)
	f()
}
