package main

type Stack[T any] struct {
	elements []T
}

func NewStack[T any]() *Stack[T] {
	return &Stack[T]{make([]T, 0)}
}

func (s *Stack[T]) Push(element T) {
	s.elements = append(s.elements, element)
}
func (s *Stack[T]) Pop() (T, bool) {
	if len(s.elements) == 0 {
		var defaultVal T
		return defaultVal, false
	}
	// 가장 마지막에 있는 element를 빼낸다.
	element := s.elements[len(s.elements)-1]
	// 마지막 element를 제외한 나머지 element를 다시 넣는다.
	s.elements = s.elements[:len(s.elements)-1]
	return element, true
}

func main() {
	stack := NewStack[int]()
	stack.Push(1)
	stack.Push(2)
	stack.Push(3)
	stack.Push(4)
	stack.Push(5)

	for {
		element, ok := stack.Pop()
		if !ok {
			break
		}
		println(element)
	}
}
