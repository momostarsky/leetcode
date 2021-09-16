package leetcode

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInsertSort(t *testing.T) {

	var nums = []int{3, 1}
	InsertSort(nums)
	assert.Equal(t, []int{1, 3}, nums)

	var n2 = []int{1, 7, 4, 6, 0, 9, 11}

	InsertSort(n2)
	assert.Equal(t, []int{0, 1, 4, 6, 7, 9, 11}, n2)

}

func Test_rotateArray(t *testing.T) {
	type args struct {
		nums  []int
		left  int
		mid   int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		{"rotateArray", args{
			nums:  []int{1, 2, 3, 4, 77, 66, 55},
			left:  0,
			mid:   4,
			right: 6,
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rotateArray(tt.args.nums, tt.args.left, tt.args.mid, tt.args.right)
			assert.Equal(t, []int{77, 66, 55, 1, 2, 3, 4}, tt.args.nums)
		})
	}
}

func TestMergeSort3(t *testing.T) {
	type args struct {
		array []int
	}
	tests := []struct {
		name string
		args args
	}{
		{"m4", args{
			array: []int{5, 9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3},
		}},
		{"m5", args{
			array: []int{9, 1, 6, 8, 14, 6, 49, 25, 4, 6, 3, 45, 67, 2, 5, 24, 56, 34, 24, 56, 2, 2, 21, 4, 1, 4, 7, 9},
		}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			MergeSort3(tt.args.array)
			fmt.Printf("%s,%v\n", tt.name, tt.args.array)
		})
	}
}
