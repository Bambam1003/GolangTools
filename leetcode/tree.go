// Package leetcode ...
package leetcode

import "fmt"

// TreeNode 二叉树节点
type TreeNode struct {
	Val   int32
	Left  *TreeNode
	Right *TreeNode
}

// preorderTraversal 前序遍历：根节点 -> 左子树 -> 右子树
func preorderTraversal(root *TreeNode) []int32 {
	var result []int32
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		result = append(result, node.Val)
		traverse(node.Left)
		traverse(node.Right)
	}
	traverse(root)
	return result
}

// inorderTraversal 中序遍历：左子树 -> 根节点 -> 右子树
func inorderTraversal(root *TreeNode) []int32 {
	var result []int32
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		result = append(result, node.Val)
		traverse(node.Right)
	}
	traverse(root)
	return result
}

// postorderTraversal 后序遍历：左子树 -> 右子树 -> 根节点
func postorderTraversal(root *TreeNode) []int32 {
	var result []int32
	var traverse func(node *TreeNode)
	traverse = func(node *TreeNode) {
		if node == nil {
			return
		}
		traverse(node.Left)
		traverse(node.Right)
		result = append(result, node.Val)
	}
	traverse(root)
	return result
}

func TestTreeNode() {
	// 创建一棵二叉树
	root := &TreeNode{Val: 1}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 3}
	root.Left.Left = &TreeNode{Val: 4}
	root.Left.Right = &TreeNode{Val: 5}

	// 前序遍历
	preorder := preorderTraversal(root)
	fmt.Println("前序遍历:", preorder)

	// 中序遍历
	inorder := inorderTraversal(root)
	fmt.Println("中序遍历:", inorder)

	// 后序遍历
	postorder := postorderTraversal(root)
	fmt.Println("后序遍历:", postorder)
}
