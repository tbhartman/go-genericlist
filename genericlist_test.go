package genericlist_test

import (
	"testing"

	"github.com/tbhartman/go-genericlist"
)

func TestList(t *testing.T) {
	glist := genericlist.New[int]()
	if glist.Front() != nil {
		t.Log("expected Front to be nil")
		t.Fail()
	}
	if glist.Back() != nil {
		t.Log("expected Back to be nil")
		t.Fail()
	}
	glist2 := genericlist.New[int]()
	glist.PushBack(0)
	glist2.PushBack(1)
	glist2.PushBack(2)
	glist.PushBackList(glist2)
	if glist.Len() != 3 {
		t.Logf("expected length to be %d; got %d", 3, glist.Len())
		t.Fail()
	}
}
