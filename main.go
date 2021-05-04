package main

import (
	"fmt"

	"github.com/nitinjangam/data-structure-go/trees"
)

func main() {
	t := &trees.Node{Key: 10}
	t.Insert(20)
	t.Insert(5)
	t.Insert(30)
	t.Insert(15)
	fmt.Println(t.Search(40))
}
