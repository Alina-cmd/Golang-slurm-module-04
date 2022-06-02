package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	slice := []int{1, 4, 6, 6, 10, 12, 12, 20}
	NodeStart := createList(slice)
	fmt.Printf("%#v\n", *NodeStart)
	NodeCurrent := NodeStart
	for NodeCurrent.Next != nil {
		fmt.Printf("%#v\n", *NodeCurrent.Next)
		NodeCurrent = NodeCurrent.Next
	}
	deleteDuplicates(NodeStart)
	NodeCurrent = NodeStart
	fmt.Printf("%#v\n", *NodeStart)
	for NodeCurrent.Next != nil {
		fmt.Printf("%#v\n", *NodeCurrent.Next)
		NodeCurrent = NodeCurrent.Next
	}
}

func createList(slice []int) *ListNode {
	node := &ListNode{
		Val:  slice[len(slice)-1],
		Next: nil,
	}
	for i := len(slice) - 2; i >= 0; i-- {
		node = &ListNode{
			Val:  slice[i],
			Next: node,
		}
	}
	return node
}

func deleteDuplicates(nodeCurrent *ListNode) {
	for nodeCurrent.Next != nil {
		if nodeCurrent.Next.Val == nodeCurrent.Val {
			nodeCurrent.Next = nodeCurrent.Next.Next
		} else {
			nodeCurrent = nodeCurrent.Next
		}
	}
}
