package BFS

import (
	"fmt"
	"testing"
)

func TestMaximumDepthOfBinaryTree(t *testing.T) {
	//root := &MaxTreeNode{
	//	Value: 2,
	//	Right: &MaxTreeNode{Value: 3, Right: &MaxTreeNode{Value: 4, Right: &MaxTreeNode{Value: 5, Right: &MaxTreeNode{Value: 6}}}},
	//}

	root := &MaxTreeNode{
		Value: 2,
		Right: &MaxTreeNode{Value: 3, Right: &MaxTreeNode{Value: 4, Right: &MaxTreeNode{Value: 5, Right: &MaxTreeNode{Value: 6}}}},
		Left:  &MaxTreeNode{Value: 9, Right: &MaxTreeNode{Value: 4}, Left: &MaxTreeNode{Value: 3}},
	}
	ans := maxDepth(root)
	fmt.Println(ans)

}
