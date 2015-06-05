package slll

import "testing"

type Num struct {
	Int int
}

func TestListTraversal(t *testing.T) {
	l := NewList()
	h := l.Head()
	h = h.PushBack(Num{Int: 1})
	h = h.PushBack(Num{Int: 2})
	h = h.PushBack(Num{Int: 3})

	// Note NewList()'s head is initialized to an empty element.
	firstElem := l.Head().Next()
	secondElem := firstElem.Next()
	thirdElem := secondElem.Next()
	fourthElem := thirdElem.Next()
	if fourthElem != nil {
		t.Error("expected nil, but received", fourthElem)
	}

	first_read := firstElem.Value.(Num)
	if first_read.Int != 1 {
		t.Error("expected 1, but received", first_read.Int)
	}
	second_read := secondElem.Value.(Num)
	if second_read.Int != 2 {
		t.Error("expected 2, but received", second_read.Int)
	}
	third_read := thirdElem.Value.(Num)
	if third_read.Int != 3 {
		t.Error("expected 3, but received", third_read.Int)
	}

}

func TestListRemoval(t *testing.T) {
	l := NewList()
	h := l.Head()
	h = h.PushBack(Num{Int: 1})
	h = h.PushBack(Num{Int: 2})
	h = h.PushBack(Num{Int: 3})

	l.RemoveHead()
	firstNum := l.Head().Value.(Num)
	if firstNum.Int != 1 {
		t.Error("expected 1, but received", firstNum.Int)
	}

	l.RemoveHead()
	firstNum = l.Head().Value.(Num)
	if firstNum.Int != 2 {
		t.Error("expected 2, but received", firstNum.Int)
	}

	l.RemoveHead()
	firstNum = l.Head().Value.(Num)
	if firstNum.Int != 3 {
		t.Error("expected 3, but received", firstNum.Int)
	}

}
