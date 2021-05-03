package stacks

import "log"

// Stack implementation
type Stack struct {
	ele []interface{}
}

//Push function is used to push element into stack
func (s *Stack) Push(el interface{}) {
	s.ele = append(s.ele, el)
}

//Pop function used to remove element from stack
func (s *Stack) Pop() interface{} {
	if len(s.ele) <= 0 {
		log.Fatalln("Cannot pop from empty stack.")
	}
	res := s.ele[len(s.ele)-1]
	s.ele = s.ele[:len(s.ele)-1]
	return res
}

//IsEmpty function to check if stack is empty
func (s *Stack) IsEmpty() bool {
	if len(s.ele) == 0 {
		return true
	}
	return false
}
