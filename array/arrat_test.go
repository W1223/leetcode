package array

import (
	"testing"
	"fmt"
)

// go test -v .\array -test.run TestMaxSubArray
func TestMaxSubArray(t *testing.T){
	nums := []int{2,1}
	v := maxSubArray(nums)
	fmt.Println(v)
}


// go test -v .\array -test.run TestGenerate
func TestGenerate(t *testing.T){
	v := generate(5)
	fmt.Println(v)
}

func TestGetRow(t *testing.T) {
	v := getRow(3)
	fmt.Println(v)
}

func TestMaxProfit(t *testing.T){
	v := maxProfit([]int{7,1,5,3,6,4})

	fmt.Println(v)
}


// [1,2,3,4,4,9,56,90]

func TestTwoSum(t *testing.T){
	v := twoSum([]int{2,7,11,15}, 18)

	fmt.Println(v)
}

// [1,2,3,4,4,9,56,90]

func TestRotate(t *testing.T){
	v:= []int{1,2,3,4,5,6,7}
	rotate(v, 3)
	fmt.Println(v)
}