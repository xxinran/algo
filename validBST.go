/**
 * Definition for a binary tree node.
 */

package main

import (
	"fmt"
	"math"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func isValidBST(root *TreeNode) bool {
	var check func(*TreeNode, int, int) bool
	check = func(node *TreeNode, lower, upper int) bool {
		if node == nil {
			return true
		}
		if node.Val < lower || node.Val > upper {
			return false
		}
		if check(node.Left, lower, node.Val) && check(node.Right, node.Val, upper) {
			return true
		}

		return false

		//replace L27-31 by return check(node.Left, lower, node.Val) && check(node.Right, node.Val, upper) also works

	}

	return check(root, math.MinInt32, math.MaxInt32)

}

func main() {
	root := &TreeNode{Val: 5}
	root.Left = &TreeNode{Val: 1}
	root.Right = &TreeNode{Val: 4}
	root.Left.Left = nil
	root.Left.Right = nil
	root.Right.Left = &TreeNode{Val: 3}
	root.Right.Right = &TreeNode{Val: 6}

	fmt.Println(isValidBST(root))
}
