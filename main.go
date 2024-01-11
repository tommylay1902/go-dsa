package main

import (
	"fmt"

	"github.com/tommylay1902/go-dsa/dynamicprogramming/canconstruct"
)

func main() {
	// tree := new(bst.BST)
	// tree.Add(5, tree.Root)
	// tree.Add(4, tree.Root)
	// tree.Add(3, tree.Root)
	// tree.Add(7, tree.Root)
	// tree.Add(6, tree.Root)

	// tree.PrintInOrderTraversal(tree.Root)
	// fmt.Println()
	// tree.Remove(6, tree.Root)

	// tree.PrintInOrderTraversal(tree.Root)

	// fmt.Println("fib result of 10:", dynamicprogramming.Fibbonacci(100))

	// fmt.Println("fib result of 10:", dynamicprogramming.FibbonacciMemo(100))

	// fmt.Println(canSum.CanSum(300, []int{7, 14}))

	// fmt.Println(howsum.HowSum(300, []int{7, 14}, make(map[int][]int)))

	// fmt.Println(bestsum.BestSum(300, []int{7, 14}))

	fmt.Println(canconstruct.CanConstructBrute("abcdef", []string{"ab", "a", "b", "c", "de"}))

}
