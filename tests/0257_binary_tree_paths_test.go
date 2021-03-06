package tests

import (
	"reflect"
	"strconv"
	"testing"
)

/**
 * [257] Binary Tree Paths
 *
 * Given a binary tree, return all root-to-leaf paths.
 *
 * Note: A leaf is a node with no children.
 *
 * Example:
 *
 *
 * Input:
 *
 *    1
 *  /   \
 * 2     3
 *  \
 *   5
 *
 * Output: ["1->2->5", "1->3"]
 *
 * Explanation: All root-to-leaf paths are: 1->2->5, 1->3
 *
 */

func TestBinaryTreePaths(t *testing.T) {
	x := &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}}
	var cases = []struct {
		input  *TreeNode
		output []string
	}{
		{
			input:  x,
			output: []string{"1->2->5", "1->3"},
		},
	}
	for _, c := range cases {
		x := binaryTreePaths(c.input)
		if !reflect.DeepEqual(x, c.output) {
			t.Fail()
		}
	}
}

// submission codes start here

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func binaryTreePaths(root *TreeNode) []string {
	var res []string
	if root == nil {
		return nil
	}
	var path string
	dfsBinaryPaths(root, path, &res)
	return res
}

func dfsBinaryPaths(root *TreeNode, path string, res *[]string) {
	if root == nil {
		return
	}
	if path == "" {
		path += strconv.Itoa(root.Val)
	} else {
		path += "->" + strconv.Itoa(root.Val)
	}
	if root.Left == nil && root.Right == nil {
		*res = append(*res, path)
	}
	dfsBinaryPaths(root.Left, path, res)
	dfsBinaryPaths(root.Right, path, res)
}

// submission codes end
