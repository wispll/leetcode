package main

import (
	"flag"
	"fmt"
	"strconv"
)

type arrayInt []int

func (a *arrayInt) String() string {
	return "my array int"
}

func (a *arrayInt) Set(value string) error {
	i, err := strconv.Atoi(value)
	if err != nil {
		return err
	}
	*a = append(*a, i)
	return nil
}

func main() {
	var nums arrayInt
	flag.Var(&nums, "l", "")
	flag.Parse()
    fmt.Printf("array : %v\n", nums)
}

type TreeNode struct {
    Val int
    Left *TreeNode
    Right *TreeNode
}


func inorderTraversal(root *TreeNode) []int {
    res := []int{}

    var inorder func(root *TreeNode)
    inorder = func(root *TreeNode){
        if root == nil {
            return 
        }
        inorder(root.Left)
        res = append(res, root.Val)
        inorder(root.Right)
    }

    inorder(root)

    return res
}
