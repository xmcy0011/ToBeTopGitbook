package benchmark

import (
	"fmt"
	"github.com/go-playground/assert/v2"
	"testing"
)

func benchFib(b *testing.B, num int) {
	for n := 0; n < b.N; n++ {
		//FibIterator(num)
		Fib(num)
	}
}
func BenchmarkFib_10(b *testing.B) {
	benchFib(b, 10)
}
func BenchmarkFib_30(b *testing.B) {
	benchFib(b, 30)
}
func BenchmarkFib_40(b *testing.B) {
	benchFib(b, 40)
}
func BenchmarkFib_Table(b *testing.B) {
	nums := []int{10, 20, 30}
	for i := 0; i < len(nums); i++ {
		num := nums[i]
		name := fmt.Sprintf("%d", num)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				FibIterator(num)
				//Fib(num)
			}
		})
	}
}
func BenchmarkFibV2_Table(b *testing.B) {
	maxNum := 30
	for num := 1; num <= maxNum; num++ {
		name := fmt.Sprintf("%s_%d", "BenchmarkFibV2_Table", num)
		b.Run(name, func(b *testing.B) {
			for n := 0; n < b.N; n++ {
				FibIterator(num)
			}
		})
	}
}

func TestFib(t *testing.T) {
	num := Fib(30)
	assert.Equal(t, 832040, num)
}
func TestFibIterator(t *testing.T) {
	num := FibIterator(30)
	assert.Equal(t, 832040, num)
}

func BenchmarkFib(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Fib(30)
	}
}
func BenchmarkFibV2(b *testing.B) {
	for n := 0; n < b.N; n++ {
		FibIterator(30) // 重复运行Fib(30)函数b.N次
	}
}

func BenchmarkFibParallel(b *testing.B) {
	i := 0
	b.RunParallel(func(p *testing.PB) {
		for p.Next() {
			_ = Fib(20) // 重复运行Fib(30)函数b.N次
			i++
		}
	})
	b.Log("BenchmarkFib,b.N=", b.N, ",RunCounts=", i)
}
func BenchmarkFib_BN_BParallel(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fib(30)
	}

	//b.RunParallel(func(pb *testing.PB) {
	//	for pb.Next() {
	//		Fib(30)
	//	}
	//})
}
