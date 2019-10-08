package tree

import "fmt"

type TreeNode struct {
	Value       int
	Left, Right *TreeNode
}

func CreateNode(Value int) *TreeNode {
	return &TreeNode{Value: Value}
}

func (node TreeNode) Print() {
	fmt.Print(node.Value)
}

func (node TreeNode) SetValue(Value int) {
	node.Value = Value
}

func (node *TreeNode) SetValue1(Value int) {
	node.Value = Value
}

func (node *TreeNode) Traverse() {
	if node == nil {
		return
	}
	node.Left.Traverse()
	node.Print()
	node.Right.Traverse()
}
