// Реализуация двусвязного списка вместе с базовыми операциями.
package list

import (
	"fmt"
)

// List - двусвязный список.
type List struct {
	root *Elem
}

// Elem - элемент списка.
type Elem struct {
	Val        interface{}
	next, prev *Elem
}

// New создаёт список и возвращает указатель на него.
func New() *List {
	var l List
	l.root = &Elem{}
	l.root.next = l.root
	l.root.prev = l.root
	return &l
}

// Push вставляет элемент в начало списка.
func (l *List) Push(e Elem) *Elem {
	e.prev = l.root
	e.next = l.root.next
	l.root.next = &e
	if e.next != l.root {
		e.next.prev = &e
	}
	return &e
}

// String реализует интерфейс fmt.Stringer представляя список в виде строки.
func (l *List) String() string {
	el := l.root.next
	var s string
	for el != l.root {
		s += fmt.Sprintf("%v ", el.Val)
		el = el.next
	}
	if len(s) > 0 {
		s = s[:len(s)-1]
	}
	return s
}

// Pop удаляет первый элемент списка.
func (l *List) Pop() *List {
	next := l.root.next
	next.prev = l.root
	l.root.next = next.next
	return l
}

// Reverse разворачивает список.
func (l *List) Reverse() *List {
	return l.reverse(l.root.next)
}

// Рекурсивная ф-ю для разворота списка.
func (l *List) reverse(e *Elem) *List {
	if e.next == l.root {
		root := e.next
		root.next, root.prev = e, root
		e.next, e.prev = e.prev, e.next
		return l
	}
	e.next, e.prev = e.prev, e.next
	return l.reverse(e.prev)
}
