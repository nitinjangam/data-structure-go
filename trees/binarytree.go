package trees

import "fmt"

//Node is implementation of tree
type Node struct {
	Key   int
	Left  *Node
	Right *Node
}

//Insert function to insert into tree
func (n *Node) Insert(k int) {
	if n.Key < k {
		if n.Right == nil {
			n.Right = &Node{Key: k}
		} else {
			n.Right.Insert(k)
		}
	} else if n.Key > k {
		if n.Left == nil {
			n.Left = &Node{Key: k}
		} else {
			n.Left.Insert(k)
		}
	}
}

//Search function to serach element in tree
func (n *Node) Search(k int) bool {

	if n.Key < k {
		if n.Right == nil {
			fmt.Println(n.Right == nil, k)
			return false
		}
		return n.Right.Search(k)
	} else if n.Key > k {
		if n.Left == nil {
			return false
		}
		return n.Left.Search(k)
	}
	return true
}
