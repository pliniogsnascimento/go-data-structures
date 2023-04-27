package ds

import "fmt"

type LinkedList[Type any] struct {
	head *Node[Type]
	Len  int
}

type Node[Type any] struct {
	data Type
	next *Node[Type]
}

func (l *LinkedList[Type]) Add(element Type) {
	if l.head == nil {
		l.head = &Node[Type]{data: element}
		l.Len++
		return
	}
	l.head = &Node[Type]{
		data: element,
		next: l.head,
	}
	l.Len++
}

func (l *LinkedList[Type]) AddIndex(element Type, index int) error {
	temp := l.head
	if index == 0 {
		l.Add(element)
		return nil
	}
	if index > l.Len {
		return fmt.Errorf("Index out of bounds")
	}
	for i := 0; i < (index - 1); i++ {
		temp = temp.next
	}
	temp.next = &Node[Type]{
		data: element,
		next: temp.next,
	}
	l.Len++
	return nil
}

func (l *LinkedList[Type]) AddEnd(element Type) {
	l.AddIndex(element, l.Len)
}

func (l *LinkedList[Type]) String() string {
	s := ""
	if l.head == nil {
		return s
	}
	temp := l.head
	for temp.next != nil {
		s = s + fmt.Sprint(temp.data, " ")
		temp = temp.next
	}
	s = s + fmt.Sprint(temp.data)
	return fmt.Sprintf("[%s]", s)
}
