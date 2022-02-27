package list_node

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func GenerateNodesLink(nums []int) *ListNode {
	if len(nums) == 0 {
		return nil
	}
	dummy := new(ListNode)
	var cur *ListNode
	n := len(nums)
	for i := 0; i < n; i++ {
		node := new(ListNode)
		node.Val = nums[i]

		if cur == nil {
			dummy.Next = node
		} else {
			cur.Next = node
		}
		cur = node
	}
	return dummy.Next
}

func PrintNode(head *ListNode) {
	fmt.Printf("print node now\n")
	for tmp := head; tmp != nil; tmp = tmp.Next {
		fmt.Printf("%v ", tmp.Val)
	}
	fmt.Printf("\n")
}

func printNodes(nodes []*ListNode) {
	fmt.Printf("print nodes now\n")
	n := len(nodes)
	for i := 0 ; i< n;i++ {
		fmt.Printf("%v ", nodes[i].Val)
	}
	fmt.Printf("\nprint nodes over\n")
}

func LinkNodes(lists []*ListNode) (head *ListNode, tail *ListNode){
	if len(lists) == 0 {
		return nil,nil
	}
	cur := lists[0]
	n := len(lists)
	for i := 1;i < n; i++ {
		cur.Next = lists[i]
	}
	return lists[0], lists[n-1]
}