package leetcode

import (
	"fmt"
	"sort"
)

/*
颜色分类
给定一个包含红色、白色和蓝色，一共 n 个元素的数组，原地对它们进行排序，使得相同颜色的元素相邻，并按照红色、白色、蓝色顺序排列。

此题中，我们使用整数 0、 1 和 2 分别表示红色、白色和蓝色。
进阶：
   你可以不使用代码库中的排序函数来解决这道题吗？
   你能想出一个仅使用常数空间的一趟扫描算法吗？

*/
func SortColors(nums []int) {
	var ln = len(nums)
	if ln <= 1 {
		return
	}

	for i := 0; i < ln; i++ {
		for j := i + 1; j < ln; j++ {
			if nums[i] > nums[j] {
				nums[i], nums[j] = nums[j], nums[i]
			}
		}
	}

}

func SortColorsWithScan(nums []int) {
	var ln = len(nums)
	if ln <= 1 {
		return
	}
	var sscaner = make(map[int]int)
	for _, v := range nums {
		if vv, mok := sscaner[v]; !mok {
			sscaner[v] = 1
		} else {
			sscaner[v] = vv + 1
		}
	}
	fmt.Printf("%v\n", sscaner)
	var keys []int
	for k := range sscaner {
		keys = append(keys, k)
	}
	sort.Ints(keys)

	var pbose = 0
	for i := 0; i < len(keys); i++ {

		skey := keys[i]
		var svalue, _ = sscaner[skey]
		for j := pbose; j < pbose+svalue; j++ {

			nums[j] = skey
		}

		pbose = pbose + svalue
	}
}

/*
前 K 个高频元素
给你一个整数数组 nums 和一个整数 k ，请你返回其中出现频率前 k 高的元素。你可以按 任意顺序 返回答案。

示例 1:

输入: nums = [1,1,1,2,2,3], k = 2
输出: [1,2]
示例 2:

输入: nums = [1], k = 1
输出: [1]


提示：

1 <= nums.length <= 105
k 的取值范围是 [1, 数组中不相同的元素的个数]
题目数据保证答案唯一，换句话说，数组中前 k 个高频元素的集合是唯一的
进阶：你所设计算法的时间复杂度 必须 优于 O(n log n) ，其中 n 是数组大小。
*/

type cpair struct {
	K int
	F int
}

func TopKFrequent(nums []int, k int) []int {

	var ln = len(nums)
	if ln <= 1 {
		return nums
	}
	var sscaner = make(map[int]int)
	for _, v := range nums {
		if vv, mok := sscaner[v]; !mok {
			sscaner[v] = 1
		} else {
			sscaner[v] = vv + 1
		}
	}
	var topx []cpair = nil
	for k, v := range sscaner {
		topx = append(topx, cpair{
			K: k,
			F: v,
		})
	}

	fmt.Printf("Befor : %d,%v\n", len(topx), topx)
	sort.Slice(topx, func(i, j int) bool {
		return topx[i].F > topx[j].F
	})
	fmt.Printf("Afert : %d,%v\n", len(topx), topx)
	var res []int = nil
	for i := 0; i < k && i < len(topx); i++ {

		res = append(res, topx[i].K)
	}
	return res

}
func FindKthLargest(nums []int, k int) int {
	sort.Ints(nums)
	// fmt.Printf("Afert : %d,%v\n", len(nums), nums)
	var res []int = nil
	var ln = len(nums)

	// res = append(res, nums[ln-1])
	// for i := ln - 1; i >= 0; i-- {
	// 	if nums[i] != res[len(res)-1] {
	// 		res = append( res, nums[i])
	// 	}
	// }
	// fmt.Printf("res : %d,%v\n", len(res), res)

	//  for j := len(res) - 1; j >= 0; j-- {
	// 	if k == 1 {
	// 		return res[j]
	// 	}
	// 	k = k - 1
	// }

	res = append(res, nums[0])

	for i := 0; i < ln; i++ {

		if nums[i] != res[len(res)-1] {
			res = append(res, nums[i])
		}

	}

	return res[k-1]
}
