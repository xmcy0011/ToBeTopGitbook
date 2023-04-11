package benchmark

// Fib 递归版本的解法，O(n^2)时间复杂度，性能很烂
func Fib(n int) int {
	if n == 0 || n == 1 {
		return n
	}
	return Fib(n-2) + Fib(n-1)
}

// FibIterator 迭代版本解法，O(n)时间复杂度，算法性能大幅度提升
func FibIterator(n int) int {
	if n <= 1 {
		return n
	}

	// 空间换时间，把前一个结果缓存起来，避免重复计算
	var n2, n1 = 0, 1
	for i := 2; i < n; i++ {
		n2, n1 = n1, n1+n2
	}

	return n2 + n1
}
