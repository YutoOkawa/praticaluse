package main

import "sync"

type BigStruct struct {
	Member string
}

var pool = &sync.Pool{
	New: func() interface{} {
		return &BigStruct{}
	},
}

func main() {
	b := pool.Get().(*BigStruct)

	pool.Put(b)
}

func NewBigStruct() *BigStruct {

	b := pool.Get().(*BigStruct)
	return b
}

func (b *BigStruct) Release() {
	b.Member = ""
	pool.Put(b)
}
