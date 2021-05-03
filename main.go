package main

import (
	"fmt"

	"github.com/nitinjangam/data-structure-go/linkedlists"
)

func main() {
	l := linkedlists.Node{
		Val: 0,
	}
	l1 := linkedlists.Node{
		Val: 1,
	}
	l2 := linkedlists.Node{
		Val: 2,
	}
	l3 := linkedlists.Node{
		Val: 3,
	}
	l4 := linkedlists.Node{
		Val: 4,
	}
	l.Next = &l1
	l1.Next = &l2
	l2.Next = &l3
	l3.Next = &l4
	//l4.Next = &l
	fmt.Println(l.Next.Val)
	fmt.Println(l1.Next.Val)
	fmt.Println(l2.Next.Val)
	fmt.Println(l3.Next.Val)
	l.ReverseList()
	fmt.Println(l.Next)
	fmt.Println(l1.Next.Val)
	fmt.Println(l2.Next.Val)
	fmt.Println(l3.Next.Val)
	fmt.Println(l4.Next.Val)
}
