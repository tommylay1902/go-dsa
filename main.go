package main

import (
	"fmt"

	"github.com/tommylay1902/go-dsa/dynamicprogramming/countconstruct"
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

	// fmt.Println("fib result of 10:", fib.Fibbonacci(100))

	// fmt.Println("fib result of 10:", fib.FibbonacciMemo(100))

	// fmt.Println(cansum.CanSum(300, []int{7, 14}))

	// fmt.Println(howsum.HowSum(300, []int{7, 14}, make(map[int][]int)))

	// fmt.Println(bestsum.BestSum(300, []int{7, 14}))
	// fmt.Println(canconstruct.CanConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"ab", "a", "b", "c", "de", "a", "a", "a", "e", "ee", "eeeeeee", "eeeeeeeeeeeeeee", "eeeeee", "eeeeeee"}))
	// fmt.Println(canconstruct.CanConstructBrute("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"ab", "a", "b", "c", "de", "a", "a", "a", "e", "ee", "eeeeeee", "eeeeeeeeeeeeeee", "eeeeee", "eeeeeee", "eee", "eeee", "eeeeeeeeeeee"}))

	fmt.Println(countconstruct.CountConstructBrute("a", []string{"ab", "a", "b", "c", "de"}))

	// fmt.Println(countconstruct.CountConstruct("eeeeeeeeeeeeeeeeeeeeeeeeeeeeeeef", []string{"ab", "a", "b", "c", "de", "a", "a", "a", "e", "ee", "eeeeeee", "eeeeeeeeeeeeeee", "eeeeee", "eeeeeee", "eee", "eeee", "eeeeeeeeeeee", "f"}))
}
