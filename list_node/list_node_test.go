package list_node

import "testing"

func TestGenerateNodesLink(t *testing.T) {
	nums := []int{1, 2, 3}
	link := GenerateNodesLink(nums)
	PrintNode(link)
}
