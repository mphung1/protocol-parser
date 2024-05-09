package structures

import "project/types"

/*
● During processing, the protocol may call for the creation of one or
more stacks
● Each stack functions as a traditional stack: last in, first out
● Stacks are identified by a single, lowercase character (e.g. ‘a’, ‘b’,
‘z’)
● There is no explicit stack creation command; stacks are created as
needed when something is pushed to it
● Stack values may be removed as part of an operation, but a stack
itself is not deleted once it is created
*/

type StackWrapper struct {
	types.Stack
}

// Clears the stack
func (s *StackWrapper) Clear() {
	s.Arr = []int{};
}

// Gets the stack arr
func (s *StackWrapper) GetStack() []int {
	return s.Arr;
}

// Checks if the stack is empty
func (s *StackWrapper) IsEmpty() bool {
	return len(s.Arr) == 0;
}

// Pops the last element from the stack
func (s *StackWrapper) Pop() int {
	val := s.Arr[len(s.Arr)-1];
	s.Arr = s.Arr[:len(s.Arr)-1];
	return val;
}

// Pushes a value to the stack
func (s *StackWrapper) Push(val int) {
	s.Arr = append(s.Arr, val);
}