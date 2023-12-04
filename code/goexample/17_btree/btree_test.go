package btree

import (
	"testing"

	"github.com/google/btree"
)

func TestLessInt(t *testing.T) {
	tr := btree.New(2)
	for i := 0; i < 10; i++ {
		tr.ReplaceOrInsert(NewLessInt(i))
	}

	var got []btree.Item
	tr.AscendRange(NewLessInt(0), NewLessInt(10), func(a btree.Item) bool {
		got = append(got, a)
		return true
	})

	for i := range got {
		println(i)
	}
}

func TestLessPath(t *testing.T) {
	tr := btree.New(2)

	// test path
	// 	|- A1
	// 		|- B1
	//		|- B2
	//      |- B3
	// 		|- 1.txt
	// 	|- A2
	//		|- B4
	//		|- B5
	//		|- 2.txt
	paths := []string{"/test", "/test/A1", "/test/A2", "/test/A1/B1", "/test/A1/B2", "/test/A1/B3", "/test/A1/1.txt",
		"/test/A2", "/test/A2/B4", "/test/A2/B5", "/test/A2/2.txt"}

	for _, v := range paths {
		tr.ReplaceOrInsert(NewLessPath(v))
	}

	var got []btree.Item
	begin := NewLessPath("/test")
	end := NewLessPath("/test/A3")
	tr.AscendRange(begin, end, func(a btree.Item) bool {
		got = append(got, a)
		return true
	})

	tr.Get(begin)

	for _, v := range got {
		println(v.(*LessPath).value)
	}
}
