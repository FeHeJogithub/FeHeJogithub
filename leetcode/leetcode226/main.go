package main

import "fmt"

// TreeNode definition (igual que LeetCode)
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// ---------------------------
// INVERTIR ÁRBOL
// ---------------------------
func invertTree(root *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	root.Left, root.Right = root.Right, root.Left

	invertTree(root.Left)
	invertTree(root.Right)

	return root
}

// ---------------------------
// PRINT (preorder) para ver el árbol
// ---------------------------
func printTree(root *TreeNode) {
	if root == nil {
		fmt.Print("nil ")
		return
	}
	fmt.Printf("%d ", root.Val)
	printTree(root.Left)
	printTree(root.Right)
}

// ---------------------------
// MAIN
// ---------------------------
func main() {
	// Construimos este árbol:
	//        4
	//      /   \
	//     2     7
	//    / \   / \
	//   1  3  6  9

	root := &TreeNode{Val: 4}
	root.Left = &TreeNode{Val: 2}
	root.Right = &TreeNode{Val: 7}
	root.Left.Left = &TreeNode{Val: 1}
	root.Left.Right = &TreeNode{Val: 3}
	root.Right.Left = &TreeNode{Val: 6}
	root.Right.Right = &TreeNode{Val: 9}

	fmt.Print("Árbol original (preorder): ")
	printTree(root)
	fmt.Println()

	invertTree(root)

	fmt.Print("Árbol invertido (preorder): ")
	printTree(root)
	fmt.Println()
}
