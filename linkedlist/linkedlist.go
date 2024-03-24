package main

import (
	"fmt"
)

type Node[T any] struct {
	value      T
	prev, next *Node[T]
}

type DoublyLinkedList[T any] struct {
	size int
	head *Node[T]
	tail *Node[T]
}

func (n *Node[T]) NewNode(v T) *Node[T] {
	return &Node[T]{
		value: v,
		prev:  nil,
		next:  nil,
	}
}
func (d *DoublyLinkedList[T]) InsertAt(index int, v T) error {
	if index < 0 || index > d.size {
		return fmt.Errorf("invalid index")
	}
	n := &Node[T]{}
	newNode := n.NewNode(v)

	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else if index == 0 {
		newNode.next = d.head
		d.head.prev = newNode
		d.head = newNode
	} else if index == d.size {
		newNode.prev = d.tail
		d.tail.next = newNode
		d.tail = newNode
	} else {
		current := d.head
		for i := 1; i < index; i++ {
			current = current.next
		}
		newNode.prev = current
		newNode.next = current.next
		current.next.prev = newNode
		current.next = newNode
	}
	d.size++
	return nil
}

func (d *DoublyLinkedList[T]) Append(v T) {
	n := &Node[T]{}
	newNode := n.NewNode(v)
	if d.head == nil {
		d.head = newNode
		d.tail = newNode
	} else {
		d.tail.next = newNode
		newNode.prev = d.tail
		d.tail = newNode
	}
	d.size++
}

func (d *DoublyLinkedList[T]) Get(index int) (*Node[T], error) {
	if d.size == 0 {
		return nil, fmt.Errorf("empty linkedlist")
	}
	if index > d.size || index < 0 {
		return nil, fmt.Errorf("invalid index")
	}

	current := d.head
	for i := 0; i < index; i++ {
		current = current.next
	}
	return current, nil
}

func (l *DoublyLinkedList[T]) AddElements(elements []struct {
	index int
	value T
}) error {
	for _, e := range elements {
		if err := l.InsertAt(e.index, e.value); err != nil {
			return err
		}
	}

	return nil
}

func (l *DoublyLinkedList[T]) PrintForward() string {
	if l.size == 0 {
		return ""
	}
	current := l.head
	output := "HEAD"
	for current != nil {
		output = fmt.Sprintf("%s -> %v", output, current.value)
		current = current.next
	}

	return fmt.Sprintf("%s -> NULL", output)
}

func (l *DoublyLinkedList[T]) PrintReverse() string {
	if l.size == 0 {
		return ""
	}
	current := l.tail
	output := "NULL"
	for current != nil {
		output = fmt.Sprintf("%s <- %v", output, current.value)
		current = current.prev
	}
	return fmt.Sprintf("%s <- HEAD", output)
}

func main() {
	dl := &DoublyLinkedList[string]{}
	//d.Append("thiri")
	//d.Append("lwin")

	dl.InsertAt(0, "C")
	dl.InsertAt(0, "A")
	dl.InsertAt(1, "B")
	// dl.InsertAt(3, "D")

	d, err := dl.Get(1)
	if err != nil {
		fmt.Printf("Error :%s", err.Error())
		return
	}
	fmt.Println("data >>", *d)
}
