package leetcode

import (
	"fmt"
)

//  https://geekr.dev/posts/go-sorting-algorithm-quick

// 快速排序入口函数
func quickSort(nums []int, p int, r int) {
	// 递归终止条件
	if p >= r {
		return
	}
	fmt.Printf("%v\n", nums)
	// 获取分区位置
	q := partition(nums, p, r)
	// 递归分区（排序是在定位 pivot 的过程中实现的）
	quickSort(nums, p, q-1)
	quickSort(nums, q+1, r)
}

// 定位 pivot
func partition(nums []int, left int, right int) int {
	// 以当前数据序列最后一个元素作为初始 pivot
	pivot := nums[right]
	// 初始化 i、j 下标
	i := left
	// 后移 j 下标的遍历过程
	for j := left; j < right; j++ {
		// 将比 pivot 小的数丢到 [p...i-1] 中，剩下的 [i...j] 区间都是比 pivot 大的
		if nums[j] < pivot {
			// 互换 i、j 下标对应数据
			nums[i], nums[j] = nums[j], nums[i]

			fmt.Printf("%v\n", nums)
			// 将 i 下标后移一位
			i++
		}
	}
	// 最后将 pivot 与 i 下标对应数据值互换
	// 这样一来，pivot 就位于当前数据序列中间，i 也就是 pivot 值对应的下标
	nums[i], nums[right] = pivot, nums[i]
	// 返回 i 作为 pivot 分区位置
	return i
}

type qkRange struct {
	start int
	end   int
}

func QuickSortIt(arr []int) {

	N := len(arr)
	if N <= 1 {
		return
	}

	//r[]模拟堆疊,p为数量,r[p++]为push,r[--p]为pop且取得元素
	var r []qkRange = nil
	r = append(r, qkRange{start: 0, end: N - 1})

	for len(r) > 0 {
		var crange = r[0]

		r = r[1:]

		fmt.Printf("stack Leng :%d\n", len(r))
		if crange.start >= crange.end {
			continue
		}
		var mid = arr[crange.end]
		var left = crange.start
		var right = crange.end - 1
		for left < right {
			for ; arr[left] < mid && left < right; left++ {
			}

			for ; arr[right] >= mid && left < right; right-- {
			}
			arr[left], arr[right] = arr[right], arr[left]
		}
		if arr[left] >= arr[crange.end] {
			arr[left], arr[crange.end] = arr[crange.end], arr[left]
		} else {
			left = left + 1
		}

		if crange.start < (left - 1) {
			r = append(r, qkRange{start: crange.start, end: left - 1})
		}
		if (left + 1) < crange.end {
			r = append(r, qkRange{start: left + 1, end: crange.end})
		}

	}
}

func QuickSort(nums []int) {

	N := len(nums)
	if N <= 1 {
		return
	}
	quickSort(nums, 0, N-1)
}

func UsHappy(n int) bool {

	var visted = make(map[int]bool)
	for n != 1 {
		var sum = 0
		var temp = n
		for temp > 0 {
			var mod = temp % 10
			temp = temp / 10
			sum = sum + mod*mod
		}

		if _, ok := visted[sum]; !ok {
			visted[sum] = true

		} else {
			return false
		}

		n = sum

	}
	return true

}
func titleToNumber(columnTitle string) int {
	// 1 <= columnTitle.length <= 7
	// columnTitle 仅由大写英文组成
	// columnTitle 在范围 ["A", "FXSHRXW"] 内

	// 作者：力扣 (LeetCode)
	// 链接：https://leetcode-cn.com/leetbook/read/top-interview-questions-medium/xweb76/
	// 来源：力扣（LeetCode）
	// 著作权归作者所有。商业转载请联系作者获得授权，非商业转载请注明出处。

	var ans = 0
	var bytes = []byte(columnTitle)
	for _, v := range bytes {
		fmt.Printf("%d", v-64)
		var bv = int(v - 64)

		ans = ans + bv*26
	}

	return ans

}
