package list

import (
	"errors"
	"fmt"
	"strconv"
)

type ListNode struct {
	Value int
	Next  *ListNode
}

func NewListNode(val int) *ListNode {
	node := ListNode{Value: val, Next: nil}
	return &node
}

func NewListNodeWithNext(val int, next *ListNode) *ListNode {
	node := ListNode{Value: val, Next: next}
	return &node
}

func AddNode(val int, head *ListNode) error {
	if head.Next != nil {
		return errors.New("Can only add to last element of list")
	}
	head.Next = &ListNode{Value: val, Next: nil}
	return nil
}

func ReverseList(head *ListNode) *ListNode {
	return reverseListHelper(head, nil)
}

func reverseListHelper(curNode *ListNode, prevNode *ListNode) *ListNode {
	if curNode == nil {
		return prevNode
	}

	var listNode *ListNode
	if prevNode == nil {
		listNode = NewListNode(curNode.Value)
	} else {
		listNode = NewListNodeWithNext(curNode.Value, prevNode)
	}

	return reverseListHelper(curNode.Next, listNode)

}

func PrintList(head *ListNode) {
	for head != nil {
		fmt.Print(strconv.Itoa(head.Value) + "->")
		head = head.Next
	}
	fmt.Println("EOL")
}
