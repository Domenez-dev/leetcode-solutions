/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	if root == nil {
		return []string{}
	}

	var result []string

	if root.Left == nil && root.Right == nil {
		return []string{fmt.Sprintf("%v", root.Val)}
	}

	leftPaths := binaryTreePaths(root.Left)
	rightPaths := binaryTreePaths(root.Right)

	for _, path := range leftPaths {
		result = append(result, fmt.Sprintf("%v->%v", root.Val, path))
	}
	for _, path := range rightPaths {
		result = append(result, fmt.Sprintf("%v->%v", root.Val, path))
	}

	return result
}
