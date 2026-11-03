package main

import(
	"goImpl/array"
	"fmt"
)

func main(){
	var myArr []int = []int{1, 2, 3, 4}
	var target int = 5
	var result []int = array.TwoSum(myArr, target)
	fmt.Printf("This is the result: %v", result)
}

