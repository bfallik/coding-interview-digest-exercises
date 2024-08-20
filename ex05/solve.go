package ex05

import (
	"fmt"
)

type Node struct {
	data int
	next *Node
}

type List struct {
	head *Node
}

func (l *List) PushBack(data int) {
	newNode := &Node{data, nil}
	if l.head == nil {
		l.head = newNode
		return
	}
	n := l.head
	for n.next != nil {
		n = n.next
	}
	n.next = newNode
}

func (l *List) AddBacklink(pos int) {
	dest := l.head
	for i := 0; i < pos; i++ {
		dest = dest.next
	}

	tail := l.head
	for tail.next != nil {
		tail = tail.next
	}

	tail.next = dest
}

func Solve(l List) string {
	var p1, p2 *Node = l.head, l.head

	for {
		if p2.next == nil {
			return "null"
		}
		p2 = p2.next
		if p2.next == nil {
			return "null"
		}

		p1 = p1.next

		if p2 == p1 { // found a loop
			history := map[*Node]int{}
			for n := l.head; n.next != nil; n = n.next {
				history[n] = len(history)
				if pos, ok := history[n.next]; ok {
					return fmt.Sprintf("index %d", pos)
				}
			}
		}
	}

}
