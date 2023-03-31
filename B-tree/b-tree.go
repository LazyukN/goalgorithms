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

func (n *TreeNode) Delete(item int) bool {
	if n.item == item {
		if n.left != nil && n.right != nil {
			rightNode := n.right
			mostLeftNode := rightNode.left
			for mostLeftNode != nil {
				if mostLeftNode.left == nil {
					rightNode = mostLeftNode
					break
				}
				rightNode = mostLeftNode.left
				mostLeftNode = rightNode.left
			}
			//n.parent.left = nil
			newItem := rightNode.item
			n.Delete(rightNode.item)
			n.item = newItem
			return true
		} else if n.left != nil {
			n.item = n.left.item
			n.left = nil
			return true
		} else if n.right != nil {
			n.item = n.right.item
			n.right = nil
		} else {
			if n.parent != nil {
				if n.parent.left == n {
					n.parent.left = nil
				}
				if n.parent.right == n {
					n.parent.right = nil
				}
			} else {
				emptyNode := TreeNode{}
				n.item = emptyNode.item
			}
		}
	} else if n.item > item {
		if n.left != nil {
			return n.left.Delete(item)
		}
	} else {
		if n.right != nil {
			return n.right.Delete(item)
		}
	}

	return false
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
	test := []int{3, 15, 16, 7, 8, 4, 5, 2}
	node := TreeNode{
		item: 12,
	}
	for _, val := range test {
		node.Insert(val)
	}
	deleted := node.Delete(2)

	fmt.Println(deleted)
	test2 := node.Traverse()
	fmt.Println(test2)
}
