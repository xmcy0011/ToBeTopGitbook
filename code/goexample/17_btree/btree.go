package btree

import "github.com/google/btree"

type LessInt struct {
	value int
}

func NewLessInt(value int) *LessInt {
	return &LessInt{value: value}
}

func (l LessInt) Less(than btree.Item) bool {
	return l.value < than.(*LessInt).value
}

type LessPath struct {
	value string
}

func NewLessPath(value string) *LessPath {
	return &LessPath{value: value}
}

func (l LessPath) Less(than btree.Item) bool {
	return l.value < than.(*LessPath).value
}
