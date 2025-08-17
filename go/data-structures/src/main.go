package main

import (
	"data-structures/src/list"
)

func main() {
	listNode := list.NewListNode(1)
	list.AddNode(2, listNode)
	list.AddNode(3, listNode.Next)
	list.PrintList(listNode)

	reverseListNode := list.ReverseList(listNode)
	list.PrintList(reverseListNode)
}
