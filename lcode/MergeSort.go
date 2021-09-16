package leetcode

// 这个算法不是自己实现的
//

import (
	"math"
)

//insertionSort  插入排序对小规模的
//数组规模 n 较小的大多数情况下，我们可以使用插入排序，它比冒泡排序，选择排序都快，甚至比任何的排序算法都快。
//数列中的有序性越高，插入排序的性能越高，因为待排序数组有序性越高，插入排序比较的次数越少。
func insertionSort(data []int, a, b int) {
	for i := a + 1; i <= b; i++ {
		for j := i; j > a && data[j] < data[j-1]; j-- {
			data[j], data[j-1] = data[j-1], data[j]
		}
	}
}

func InsertSort(nums []int) {
	N := len(nums)
	if N <= 1 {
		return
	}
	if N == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return
	}
	insertionSort(nums, 0, N-1)
}

//翻转算法，也叫手摇算法，主要用来对数组两部分进行位置互换，比如数组： [9,8,7,1,2,3]，将前三个元素与后面的三个元素交换位置，变成 [1,2,3,9,8,7]。
//再比如，将字符串 abcde1234567 的前 5 个字符与后面的字符交换位置，那么手摇后变成：1234567abcde。
//如何翻转呢？
//将前部分逆序
//将后部分逆序
//对整体逆序
//示例如下：
//1. 分成两部分：[abcde][1234567]
//2. 分别逆序变成：[edcba][7654321]
//3. 整体逆序：[1234567abcde]

func rotateArray(nums []int, left, mid, right int) {
	var revers = func(start, end int) {
		for i, j := start, end; i < j; i, j = i+1, j-1 {
			nums[i], nums[j] = nums[j], nums[i]
		}

	}
	revers(left, mid-1)
	revers(mid, right)
	revers(left, right)

}

// MergeSort 归并排序
func MergeSort(nums []int) []int {
	N := len(nums)
	if N <= 1 {
		return nums
	}
	if N == 2 {
		if nums[0] > nums[1] {
			nums[0], nums[1] = nums[1], nums[0]
		}
		return nums
	}

	// 获取分区位置
	p := N / 2
	// 通过递归分区
	left := MergeSort(nums[0:p])
	right := MergeSort(nums[p:])

	//fmt.Printf("pos:%d,left :%v\n", p, left)
	//fmt.Printf("pos:%d,right :%v\n", p, right)

	return merge(left, right)
}

func merge(left []int, right []int) []int {
	i, j := 0, 0
	m, n := len(left), len(right)
	// 用于存放结果集
	var result []int
	for {
		// 任何一个区间遍历完，则退出
		if i >= m || j >= n {
			break
		}
		// 对所有区间数据进行排序
		if left[i] < right[j] {
			result = append(result, left[i])
			i++
		} else {
			result = append(result, right[j])
			j++
		}
	}
	// 如果左侧区间还没有遍历完，将剩余数据放到结果集

	for ; i < m; i++ {
		result = append(result, left[i])
	}

	// 如果右侧区间还没有遍历完，将剩余数据放到结果集

	for ; j < n; j++ {
		result = append(result, right[j])
	}

	// 返回排序后的结果集
	return result
}

///////////////////////////
///-------
// //http://m.zyiz.net/tech/detail-226586.html
// 在数学上，斐波那契数是以递归的方法来定义：

//  Func(0)=0
//  Func(1) =1
//  Func(N) = Func(N-1) + Func(N-2)   N>= 2

// Fib 首几个斐波那契数是：
// 1, 1, 2, 3, 5, 8, 13, 21, 34, 55,
////////////////////////////////////////
//普通递归实现
func Fib(n int64) (res int64) {
	if n < 2 {
		return n
	} else {
		return Fib(n-1) + Fib(n-2)
	}
}

//尾递归实现
func fib2(first, second, n int64) (res int64) {
	if n == 0 {
		return 0
	} else if n == 1 || n == 2 {
		return 1
	} else if n == 3 {
		return first + second
	}
	return fib2(second, first+second, n-1)
}

func Fib2(n int64) (res int64) {

	return fib2(1, 1, n)
}

// Fib3 循环实现
func Fib3(n int64) (res int64) {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}

	var first, second int64 = 1, 1
	var ret, i int64 = 0, 0
	for i = 3; i <= n; i++ {
		ret = first + second
		first = second
		second = ret
	}
	return second
}

const (
	HalfSqrt5 = 1.1180339887498948482045868343656
	AllSqrt5  = 2.2360679774997896964091736687313
)

//四舍五入
func round(x float64) (res int64) {
	return int64(math.Floor(x + 0.5))
}

// Fib4 用通项公式实现
func Fib4(n int64) (res int64) {
	if n <= 1 {
		return n
	}

	res = round((math.Pow(0.5+HalfSqrt5, float64(n)) - math.Pow(0.5-HalfSqrt5, float64(n))) / AllSqrt5)
	return res
}

//https://goa.lenggirl.com/#/algorithm/sort/merge_sort

// 自底向上归并排序优化版本
func MergeSort3(array []int) {
	// 按照三个元素为一组进行小数组排序，使用直接插入排序]
	n := len(array)
	blockSize := 3
	a, b := 0, blockSize
	for b <= n {
		InsertSort(array[a:b])
		a = b
		b += blockSize
	}
	InsertSort(array[a:n])

	// 将这些小数组进行归并
	for blockSize < n {
		a, b = 0, 2*blockSize
		for b <= n {
			mergeOption(array, a, a+blockSize, b)
			a = b
			b += 2 * blockSize
		}
		if m := a + blockSize; m < n {
			mergeOption(array, a, m, n)
		}
		blockSize *= 2
	}
}

// 原地归并操作
func mergeOption(array []int, begin, mid, end int) {
	// 三个下标，将数组 array[begin,mid] 和 array[mid,end-1]进行原地归并
	i, j, k := begin, mid, end-1 // 因为数组下标从0开始，所以 k = end-1

	for j-i > 0 && k-j >= 0 {
		step := 0
		// 从 i 向右移动，找到第一个 array[i]>array[j]的索引
		for j-i > 0 && array[i] <= array[j] {
			i++
		}

		// 从 j 向右移动，找到第一个 array[j]>array[i]的索引
		for k-j >= 0 && array[j] <= array[i] {
			j++
			step++
		}

		// 进行手摇翻转，将 array[i,mid] 和 [mid,j-1] 进行位置互换
		// mid 是从 j 开始向右出发的，所以 mid = j-step
		merge_rotation(array, i, j-step, j-1)
		i = i + step
	}

}

// 手摇算法，将 array[l,l+1,l+2,...,mid-2,mid-1,mid,mid+1,mid+2,...,r-2,r-1,r] 从mid开始两边交换位置
// 1.先逆序前部分：array[mid-1,mid-2,...,l+2,l+1,l]
// 2.后逆序后部分：array[r,r-1,r-2,...,mid+2,mid+1,mid]
// 3.上两步完成后：array[mid-1,mid-2,...,l+2,l+1,l,r,r-1,r-2,...,mid+2,mid+1,mid]
// 4.整体逆序： array[mid,mid+1,mid+2,...,r-2,r-1,r,l,l+1,l+2,...,mid-2,mid-1]
func merge_rotation(array []int, l, mid, r int) {
	merge_reverse(array, l, mid-1)
	merge_reverse(array, mid, r)
	merge_reverse(array, l, r)
}

func merge_reverse(array []int, l, r int) {
	for l < r {
		// 左右互相交换
		array[l], array[r] = array[r], array[l]
		l++
		r--
	}
}
