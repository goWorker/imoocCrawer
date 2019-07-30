package main

import (
	"fmt"
	"imooc/tree"
)

func main() {
	var root tree.TreeNode
	root = tree.TreeNode{Value:3}
	root.Left = &tree.TreeNode{}
	root.Right = &tree.TreeNode{5,nil,nil}
	root.Right.Left = new(tree.TreeNode)
	//nodes := []TreeNode{
	//	{value:3},
	//	{},
	//	{6,nil,&root},
	//}
	//fmt.Println(nodes)
	root.Left.Right = tree.CreateNode(2)
	root.Right.Left.SetValue(4)
	root.Right.Left.Print()
	fmt.Println("\n")
	root.Right.Left.SetValue1(4)
	root.Right.Left.Print()
	fmt.Println()
	root.Traverse()
}
