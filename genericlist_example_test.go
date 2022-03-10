package genericlist_test

import (
	"fmt"

	"github.com/tbhartman/go-genericlist"
)

func ExampleNew() {
	glist := genericlist.New[int]()
	glist.PushBack(0)
	glist.PushBack(1)
	glist.PushBack(2)
	glist.PushBack(3)
	fmt.Printf("length: %d\n", glist.Len())
	fmt.Printf("first:  %d\n", glist.Front().Value)
	fmt.Printf("back:   %d\n", glist.Back().Value)
	// output:
	// length: 4
	// first:  0
	// back:   3
}
