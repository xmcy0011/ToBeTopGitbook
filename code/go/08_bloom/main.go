package main

import (
	"fmt"
	"github.com/willf/bloom"
	"strconv"
)

func main() {
	// bit位数
	n := uint(100000)

	b1 := bloom.New(n, 5)
	initTestBloom(b1, n)
	matched1 := testMatchBloom(b1, n)

	// 增加一个放大系数
	blowUp := uint(20)
	b2 := bloom.New(n*blowUp, 5)
	initTestBloom(b2, n)
	matched2 := testMatchBloom(b2, n)

	// 计算 误判率
	total := int(n) + 100
	errRate1 := float32(matched1-int(n)) / float32(total)
	errRate2 := float32(matched2-int(n)) / float32(total)

	fmt.Printf("matched:\n b1:%d,errorRate:%f  \n b2:%d,errorRate:%f \n", matched1, errRate1, matched2, errRate2)
}

func initTestBloom(b *bloom.BloomFilter, n uint) {
	// 初始化1000条数据到 bloom 过滤器中
	for i := 0; i < int(n); i++ {
		b.Add([]byte(strconv.Itoa(i)))
	}
}

func testMatchBloom(b *bloom.BloomFilter, n uint) int {
	exist := 0
	// 在原有的基础上新增100条，判断是否存在bloom中
	for j := 0; j <= (int(n) + 100); j++ {
		if b.Test([]byte(strconv.Itoa(j))) {
			exist++
		}
	}
	return exist
}
