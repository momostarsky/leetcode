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

	fmt.Println(" 原始数据 ,中文麻烦， 讲究使用还行, 必须采用 Dash to Plank ！， 这个大一点点 ", nums)
	now := time.Now()
	for _, v := range nums {
		redBlackTree.Add(v, v)
	}

	fmt.Println("zong", now.Sub(time.Now()))
	redBlackTree.PrintPreOrder()
	fmt.Println("节点数量多少:", redBlackTree.GetTreeSize())

}
