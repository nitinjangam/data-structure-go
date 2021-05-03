package linkedlists

import (
	"log"

	"github.com/nitinjangam/data-structure-go/stacks"
)

// Node for linked list
type Node struct {
	Val  int
	Next *Node
}

//IsCircle function to check if linked lists are circularly linked
func IsCircle(n Node) bool {
	start := n
	for n.Next != nil {
		if start == *n.Next {
			return true
		}
		n = *n.Next
	}
	return false
}

//ReverseList function to reverse linked lists
func (n *Node) ReverseList() {
	if IsCircle(*n) {
		log.Fatalln("Cannot reverse list as it is circularly linked..")
	}
	s := stacks.Stack{}
	for n != nil {
		s.Push(n)
		n = n.Next
	}
	//fmt.Println("Stack:", s)
	lsts := []*Node{}
	for s.IsEmpty() != true {
		p := s.Pop()
		lsts = append(lsts, p.(*Node))
	}
	//fmt.Println("List:", lsts)
	for i := 0; i < len(lsts)-1; i++ {
		lsts[i].Next = lsts[i+1]
		//fmt.Println(lsts[i], lsts[i].Next)
	}
	//fmt.Println("List:", lsts)
	lsts[len(lsts)-1].Next = nil
}

// func main() {
// 	n := Node{
// 		val:  0,
// 		Next: &Node{},
// 	}
// 	n1 := Node{
// 		val:  1,
// 		Next: &Node{},
// 	}
// 	n2 := Node{
// 		val:  2,
// 		Next: &Node{},
// 	}
// 	n3 := Node{
// 		val:  3,
// 		Next: &Node{},
// 	}
// 	n.Next = &n1
// 	n1.Next = &n2
// 	n2.Next = &n3
// 	n3.Next = &n
// 	fmt.Println(isCircle(n))
// }
