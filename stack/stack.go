package main

import "fmt"

// Generic stack implementation
type Stack[T any] struct {
	elements []T
}

// Push pushes element into stack
func (t *Stack[T]) Push(element T) {
	t.elements = append(t.elements, element)
}

// IsEmpty returns a bool to indecate if the task is empty or not
func (t *Stack[T]) IsEmpty() bool {
	return len(t.elements) == 0
}

// Pop removes the last element from the stask
func (t *Stack[T]) Pop() (*T, error) {
	if t.IsEmpty() {
		return nil, fmt.Errorf("stack is empty")
	}
	lastElement := t.elements[len(t.elements)-1]
	t.elements = t.elements[:len(t.elements)-1]
	return &lastElement, nil
}

func main() {
	fruits := Stack[string]{}
	fmt.Println(fruits.IsEmpty())

	fruits.Push("red")
	fruits.Push("orange")

	element, err := fruits.Pop()
	if err != nil {
		fmt.Printf("Error %s", err.Error())
	}
	fmt.Println("element: ", *element)
}
