package bst

import (
	"fmt"
)

type Node struct {
	value int64
	left  *Node
	right *Node
}

type BST struct {
	Root *Node
}

func (bst *BST) Add(value int64, traverse *Node) *Node {
	if traverse == nil {
		return nil
	}

	if value < traverse.value {
		traverse.left = bst.Add(value, traverse.left)
	} else if value > traverse.value {
		traverse.right = bst.Add(value, traverse.right)
	} else {
		if traverse.right == nil {
			return traverse.left
		}

		if traverse.left == nil {
			return traverse.right
		}

		s := bst.successor(traverse.right)
		traverse.value = s.value
		traverse.right = bst.Remove(s.value, traverse.right)
	}
	return traverse
}

func (bst *BST) successor(traverse *Node) *Node {

	if traverse.left == nil {
		return traverse
	}

	return bst.successor(traverse.left)
}

func (bst *BST) Remove(value int64, traverse *Node) *Node {
	if traverse == nil {
		return nil
	}

	if value < traverse.value {
		traverse.left = bst.Remove(value, traverse.left)
	} else if value > traverse.value {
		traverse.right = bst.Remove(value, traverse.right)
	} else {
		if traverse.left == nil {
			return traverse.right
		}
		if traverse.right == nil {
			return traverse.left
		}
		s := bst.successor(traverse.right)
		traverse.right = bst.Remove(s.value, s)
	}

	return traverse
}

func (bst *BST) PrintInOrderTraversal(traverse *Node) {
	if traverse == nil {
		return
	}

	bst.PrintInOrderTraversal(traverse.left)
	fmt.Println(traverse.value)
	bst.PrintInOrderTraversal(traverse.right)

}
