package main

import (
	"fmt"
)

type (
	TreeNode struct {
		item   int
		left   *TreeNode
		right  *TreeNode
		parent *TreeNode
	}
)

func (n *TreeNode) Insert(item int) bool {

	if n.item == item {
		return false
	}
	if n.item > item {
		if n.left == nil {
			n.left = &TreeNode{
				item:   item,
				parent: n,
			}
		} else {
			return n.left.Insert(item)
		}
	}

	if n.item < item {
		if n.right == nil {
			n.right = &TreeNode{
				item:   item,
				parent: n,
			}
		} else {
			return n.right.Insert(item)
		}
	}

	return true
}

func (n *TreeNode) Delete() {
	//TODO
}

func (n *TreeNode) Traverse() []int {
	items := []int{}
	if n.left != nil {
		items = append(items, n.left.Traverse()...)
	}
	items = append(items, n.item)
	if n.right != nil {
		items = append(items, n.right.Traverse()...)
	}

	return items
}

func (n *TreeNode) Find(item int) (TreeNode, bool) {
	if n.item == item {
		return *n, true
	} else if n.item > item {
		if n.left == nil {
			return TreeNode{}, false
		} else {
			return n.left.Find(item)
		}
	} else {
		if n.right == nil {
			return TreeNode{}, false
		} else {
			return n.right.Find(item)
		}
	}
}

func main() {
	test := []int{3, 15, 16, 7, 8, 4, 5, 2, 7, 7}
	node := TreeNode{
		item: 12,
	}
	for _, val := range test {
		node.Insert(val)
	}

	test2 := node.Traverse()
	fmt.Println(test2)
}
