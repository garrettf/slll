package slll

import (
	"sync/atomic"
	"unsafe"
)

type List struct {
	head unsafe.Pointer
}

type Element struct {
	Value interface{}
	next  unsafe.Pointer
}

func NewList() *List {
	return &List{head: unsafe.Pointer(new(Element))}
}

func loadElemFromPointer(e *unsafe.Pointer) *Element {
	return (*Element)(atomic.LoadPointer(e))
}

func (l *List) Head() *Element {
	return loadElemFromPointer(&l.head)
}

func (l *List) RemoveHead() {
	curHead := l.Head()
	newHead := atomic.LoadPointer(&curHead.next)
	oldHead := (*Element)(atomic.SwapPointer(&l.head, newHead))
	atomic.StorePointer(&oldHead.next, nil)
}

func (e *Element) Next() *Element {
	return loadElemFromPointer(&e.next)
}

func (e *Element) PushBack(val interface{}) *Element {
	newElem := Element{Value: val}
	newElemPtr := unsafe.Pointer(&newElem)
	for !atomic.CompareAndSwapPointer(&e.next, nil, newElemPtr) {
		e = e.Next()
	}
	return &newElem
}
