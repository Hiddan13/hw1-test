package main

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct { //Element
	Value    interface{}
	Next     *ListItem
	Prev     *ListItem
	thislist *list
}

func next(l *ListItem) *ListItem {
	if p := l.Next; l.thislist != nil && p != &l.thislist.root {
		return p
	}
	return nil
}
func prev(l *ListItem) *ListItem {
	if p := l.Prev; l.thislist != nil && p != &l.thislist.root {
		return p
	}
	return nil
}

func (l *list) Init() *list {
	l.root.Next = &l.root
	l.root.Prev = &l.root
	l.len = 0
	return l
}

type list struct {
	List // Remove me after realization. ??????????
	// Place your code here.
	len  int
	root ListItem
}

func NewList() *list {
	return new(list).Init()
}
func (l *list) Len() int {
	return l.len
}
func (l *list) Front() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.root.Next

}
func (l *list) Back() *ListItem {
	if l.len == 0 {
		return nil
	}
	return l.root.Prev
}
func (l *list) PushBackList(other *list) {
	l.lazyInit()
	for i, e := other.Len(), other.Front(); i > 0; i, e = i-1, e.Next {
		l.insertValue(e.Value, l.root.Prev)
	}
}
func (l *list) lazyInit() {
	if l.root.Next == nil {
		l.Init()
	}
}
func (l *list) insertValue(v any, at *ListItem) *ListItem {
	return l.insert(&ListItem{Value: v}, at)
}
func (l *list) insert(e, at *ListItem) *ListItem {
	e.Prev = at
	e.Next = at.Next
	e.Prev.Next = e
	e.Next.Prev = e
	e.thislist = l
	l.len++
	return e
}
func (l *list) PushFront(v any) *ListItem {
	l.lazyInit()
	return l.insertValue(v, &l.root)
}
func (l *list) PushBack(v any) *ListItem {
	l.lazyInit()
	return l.insertValue(v, l.root.Prev)
}
func (l *list) remove(e *ListItem) {
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev
	e.Next = nil // avoid memory leaks
	e.Prev = nil // avoid memory leaks
	e.thislist = nil
	l.len--
}
func (l *list) Remove(e *ListItem) any {
	if e.thislist == l {
		// if e.list == l, l must have been initialized when e was inserted
		// in l or l == nil (e is a zero Element) and l.remove will crash
		l.remove(e)
	}
	return e.Value
}
func (l *list) move(e, at *ListItem) {
	if e == at {
		return
	}
	e.Prev.Next = e.Next
	e.Next.Prev = e.Prev

	e.Prev = at
	e.Next = at.Next
	e.Prev.Next = e
	e.Next.Prev = e
}
func (l *list) MoveToFront(e *ListItem) {
	if e.thislist != l || l.root.Next == e {
		return
	}
	// see comment in List.Remove about initialization of l
	l.move(e, &l.root)
}
func main() {

}
