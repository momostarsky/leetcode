package leetcode

import (
	"fmt"
	"math/rand"
	"testing"
	"time"
)

func TestNewRedBlackTree(t *testing.T) {

	count := 10
	redBlackTree := NewRedBlackTree()
	nums := make([]int, 0)
	for i := 0; i < count; i++ {
		nums = append(nums, rand.Intn(count))
	}

	fmt.Println(" 原始数据 ,中文麻烦： ", nums)
	now := time.Now()
	for _, v := range nums {
		redBlackTree.Add(v, v)
	}

	fmt.Println("zong", now.Sub(time.Now()))
	redBlackTree.PrintPreOrder()
	fmt.Println("节点数量:", redBlackTree.GetTreeSize())

}
