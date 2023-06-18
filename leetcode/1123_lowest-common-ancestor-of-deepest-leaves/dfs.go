/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

func rec(node *TreeNode) (*TreeNode, int) {
	if node.Left == nil && node.Right == nil {
		return node, 0
	}
	if node.Right == nil {
		nodeLeft, depthLeft := rec(node.Left)
		return nodeLeft, depthLeft + 1
	}
	if node.Left == nil {
		nodeRight, depthRight := rec(node.Right)
		return nodeRight, depthRight + 1
	}
	nodeLeft, depthLeft := rec(node.Left)
	nodeRight, depthRight := rec(node.Right)
	if depthLeft < depthRight {
		return nodeRight, depthRight + 1
	}
	if depthLeft > depthRight {
		return nodeLeft, depthLeft + 1
	}
	return node, depthLeft + 1
}

func lcaDeepestLeaves(root *TreeNode) *TreeNode {
	node, _ := rec(root)
	return node
}
