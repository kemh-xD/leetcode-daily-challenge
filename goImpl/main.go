package main

import(
	myArray "goImpl/array"
	useMath "goImpl/math"
	"fmt"
)

func main(){
	var myArr []int = []int{1, 2, 3, 4}
	var target int = 5
	var result []int = myArray.TwoSum2(myArr, target)
	fmt.Printf("This is the result: %v\n", result)

	l1 := useMath.CreateLinkedList([]int{1, 5, 7})
	l2 := useMath.CreateLinkedList([]int{5, 7, 2})

	var l3 useMath.ListNode = *useMath.AddTwoNumbers(l1, l2)

	useMath.PrintList(&l3)
	

}

