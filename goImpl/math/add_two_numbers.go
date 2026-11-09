package myMath

import "fmt"

//The node has two reference : The data it stores and the reference to the next node in the sequence/list
type ListNode struct{
	Val int
	Next *ListNode
}
//Tag : Math, LinkedList and Recursive
//it is about Linked list
//So a Linked List is a datastructure consisted of a sequence of Nodes, 
// where each node is Linked to the next
//The Goal : l1[2, 4, 3] + l2[5, 6, 4] = [8, 0, 7]
//When working with LinkedList think about the HEAD, TAIL, CURRENT_VALUE, NEXT_VALUE and LOOP
//Work on purpose
func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode{
	head := &ListNode{Val: 0}
	current := head
	for l1 != nil || l2 != nil {
		sum := 0
		if l1 != nil{
			sum += l1.Val
		}

		if l2 != nil{
			sum += l2.Val
		}
		
		current.Next = &ListNode{Val: sum}
		current = current.Next

		if l1 != nil {
			l1 = l1.Next
		}

		if l2 != nil {
			l2 = l2.Next
		}
	}
	return head.Next
}

// For LinkedList creation from an array as input
func CreateLinkedList(nums []int) *ListNode{
	head := &ListNode{Val: 0} // Hold the memery Location
	current := head
	for _, num := range nums{
		current.Next = &ListNode{Val: num}
		current = current.Next
	}
	return head.Next
}

// For printing te LinkedList
func PrintList(node *ListNode){
	var result []int

	for node != nil{
		result = append(result, node.Val)
		node = node.Next
	}

	fmt.Println(result)
}

func CountNodes(l1 *ListNode) int{
	var count int = 1
	var currentNode *ListNode = l1
	for{
		if currentNode.Next == nil {
			break
		}else{
			currentNode = currentNode.Next
			count++
		}
	}
	return count
}