// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"time"

	"github.com/pliniogsnascimento/go-data-structures/ds"
)

type List[E any] struct {
	elements []E
}

func (l *List[E]) Add(element E) {
	l.elements = append(l.elements, element)
}

func (l *List[E]) Remove(element E) {
	l.elements = append(l.elements, element)
}

func timer(name string) func() {
	start := time.Now()
	return func() {
		fmt.Printf("%s took %v\n", name, time.Since(start))
	}
}

func main() {
	fmt.Println("Hello, 世界")
	numbersList := List[int64]{}
	numbersList.Add(2)
	fmt.Println(numbersList)

	linkedListAdd()
	linkedListAddEnd()
}

func linkedListAdd() {
	n1Linked := ds.LinkedList[int]{}

	fmt.Println(n1Linked.String(), n1Linked.Len)
	defer timer("LinkedList")()
	for i := 1; i <= 1e7; i++ {
		n1Linked.Add(i)
	}

	fmt.Printf("LinkedList size: %d\n", n1Linked.Len)
}

func linkedListAddEnd() {
	n1Linked := ds.LinkedList[int]{}

	fmt.Println(n1Linked.String(), n1Linked.Len)
	defer timer("LinkedListEnd")()
	for i := 1; i <= 1e7; i++ {
		n1Linked.AddEnd(i)
	}
	fmt.Printf("LinkedList size: %d\n", n1Linked.Len)
}
