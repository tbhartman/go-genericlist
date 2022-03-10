package genericlist

import (
	"container/list"
)

type Element[T any] struct {
	Value   T
	element *list.Element
}

func (e *Element[T]) Next() *Element[T] {
	ret := e.element.Next()
	if e != nil {
		return ret.Value.(*Element[T])
	}
	return nil
}
func (e *Element[T]) Prev() *Element[T] {
	ret := e.element.Next()
	if e != nil {
		return ret.Value.(*Element[T])
	}
	return nil
}

type List[T any] struct {
	list *list.List
}

func New[T any]() *List[T] {
	return &List[T]{list: list.New()}
}

func (l *List[T]) Back() *Element[T] {
	e := l.list.Back()
	if e != nil {
		return e.Value.(*Element[T])
	}
	return nil
}
func (l *List[T]) Front() *Element[T] {
	e := l.list.Front()
	if e != nil {
		return e.Value.(*Element[T])
	}
	return nil
}
func (l *List[T]) Init() *List[T] {
	l.list.Init()
	return l
}
func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	e := &Element[T]{
		Value: v,
	}
	e.element = l.list.InsertAfter(e, mark.element)
	return e
}
func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	e := &Element[T]{
		Value: v,
	}
	e.element = l.list.InsertBefore(e, mark.element)
	return e
}
func (l *List[T]) Len() int {
	return l.list.Len()
}
func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	l.list.MoveAfter(e.element, mark.element)
}
func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	l.list.MoveBefore(e.element, mark.element)
}
func (l *List[T]) MoveToBack(e *Element[T]) {
	l.list.MoveToBack(e.element)
}
func (l *List[T]) MoveToFront(e *Element[T]) {
	l.list.MoveToFront(e.element)
}
func (l *List[T]) PushBack(v T) *Element[T] {
	e := &Element[T]{
		Value: v,
	}
	e.element = l.list.PushBack(e)
	return e
}
func (l *List[T]) PushBackList(other *List[T]) {
	l.list.PushBackList(other.list)
}
func (l *List[T]) PushFront(v T) *Element[T] {
	e := &Element[T]{
		Value: v,
	}
	e.element = l.list.PushFront(e)
	return e
}
func (l *List[T]) PushFrontList(other *List[T]) {
	l.list.PushFrontList(other.list)
}
func (l *List[T]) Remove(e *Element[T]) T {
	l.list.Remove(e.element)
	return e.Value
}
