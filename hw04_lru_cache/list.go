package hw04lrucache

type List interface {
	Len() int
	Front() *ListItem
	Back() *ListItem
	PushFront(v interface{}) *ListItem
	PushBack(v interface{}) *ListItem
	Remove(i *ListItem)
	MoveToFront(i *ListItem)
}

type ListItem struct {
	Value interface{}
	Next  *ListItem
	Prev  *ListItem
}

type list struct {
	len  int
	head *ListItem
	tail *ListItem
}

func (l *list) Len() int {
	return l.len
}

func (l *list) Front() *ListItem {
	return l.head
}
func (l *list) Back() *ListItem {
	return l.tail
}

func (l *list) PushFront(v interface{}) *ListItem {
	i := ListItem{v, nil, nil}
	l.pushFront(&i)

	return &i
}

func (l *list) pushFront(i *ListItem) {
	l.len++
	if l.head == nil && l.tail == nil {
		l.head = i
		l.tail = i
	} else if l.head == l.tail {
		i.Next = l.tail
		l.head = i
		l.tail.Prev = l.head
	} else {
		i.Next = l.head
		l.head.Prev = i
		l.head = i
	}
}

func (l *list) PushBack(v interface{}) *ListItem {
	l.len++
	i := ListItem{v, nil, nil}
	if l.head == nil && l.tail == nil {
		l.head = &i
		l.tail = &i
	} else if l.head == l.tail {
		i.Prev = l.head
		l.tail = &i
		l.head.Next = l.tail
	} else {
		i.Prev = l.tail
		l.tail.Next = &i
		l.tail = &i
	}

	return &i
}

func (l *list) Remove(i *ListItem) {
	l.len--
	if l.head == l.tail {
		l.head = nil
		l.tail = nil
	} else if i.Prev == nil {
		l.head = i.Next
		l.head.Prev = nil
		i.Next = nil
	} else if i.Next == nil {
		l.tail = i.Prev
		l.tail.Next = nil
		i.Prev = nil
	} else {
		i.Prev.Next, i.Next.Prev = i.Next, i.Prev
	}
}

func (l *list) MoveToFront(i *ListItem) {
	if i != l.head {
		l.Remove(i)
		l.pushFront(i)
	}
}

func NewList() List {
	return new(list)
}
