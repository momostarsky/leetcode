package leetcode

import (
	"fmt"
	"math"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

func reverseString(s []byte) {

	lx := len(s)
	if lx <= 1 {
		return
	}
	for i := 0; i < lx/2; i++ {
		s[i], s[lx-1-i] = s[lx-1-i], s[i]
	}
}

func reverse(x int) int {

	str := fmt.Sprintf("%d", x)
	lx := len(str)

	pEnd := lx - 1
	for ; str[pEnd] == 0; pEnd-- {

	}

	var res = make([]byte, lx)
	var sc = lx

	for pStart, pEnd := 0, lx-1; pStart < pEnd; pStart, pEnd = pStart+1, pEnd-1 {

		res[sc-1] = str[pEnd]
		sc = sc - 1
	}

	v, _ := strconv.ParseInt(str, 10, 64)

	return int(v)

}

func firstUniqChar(s string) int {

	var lx = len(s)
	if lx == 0 {
		return -1
	}
	if lx == 1 {
		return 0
	}

	for pj := 0; pj < lx; pj++ {

		var c = s[pj : pj+1]
		fmt.Printf("%v, ", c)
		if strings.LastIndex(s, c) == strings.Index(s, c) {
			return pj
		}
		if false == strings.Contains(s[pj+1:], c) && false == strings.Contains(s[0:pj], c) {
			return pj
		}

	}

	return -1

}

func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	s1 := SortString(s)
	t1 := SortString(t)
	fmt.Printf("Sorted:%s\n", s1)
	fmt.Printf("Sorted:%s\n", t1)

	var hs = make(map[string]int, len(s))
	var ht = make(map[string]int, len(s))

	for pj := 0; pj < len(s); pj++ {

		if v, mok := hs[s[pj:pj+1]]; !mok {
			hs[s[pj:pj+1]] = 1
		} else {
			hs[s[pj:pj+1]] = v + 1
		}

		if r, mok := ht[t[pj:pj+1]]; !mok {
			ht[t[pj:pj+1]] = 1
		} else {
			ht[t[pj:pj+1]] = r + 1
		}

	}

	for k2 := range hs {

		if hs[k2] != ht[k2] {

			return false
		}

	}
	return true

}

func isYesLetter(c rune) bool {

	if c >= '0' && c <= '9' {
		return true
	}

	if c >= 'a' && c <= 'z' {
		return true
	}

	if c >= 'A' && c <= 'Z' {
		return true
	}

	return false

}

func isPalindrome(s string) bool {

	arr := []rune(strings.ToLower(s))
	pl := 0
	pr := len(arr) - 1

	for pl < pr {
		if !isYesLetter(arr[pl]) {
			pl = pl + 1
			continue
		}

		if !isYesLetter(arr[pr]) {
			pr = pr - 1
			continue
		}

		if arr[pl] != arr[pr] {
			fmt.Printf("%d,%d\n", pl, pr)
			return false
		} else {
			pl, pr = pl+1, pr-1
		}

	}

	return true

}
func myAtoi(s string) int {

	regx := regexp.MustCompile(`^\s*(?P<num>[+-]*\d{1,})\s*`)
	if regx == nil {
		return 0
	}
	if !regx.MatchString(s) {
		fmt.Printf("NOT Mathc")
		return 0
	}

	match := regx.FindStringSubmatch(s)
	groupNames := regx.SubexpNames()
	var numstr = ""
	for j, name := range groupNames {
		if j != 0 && name == "num" {
			numstr = match[j]
		}
	}
	if len(numstr) == 0 {
		return 0
	}

	for numstr[0:1] == "0" {
		numstr = numstr[1:]
	}

	res, _ := strconv.ParseInt(numstr, 10, 64)
	if res >= math.MaxInt32 {
		return math.MaxInt32
	} else if res <= math.MinInt32 {
		return math.MinInt32
	}
	return int(res)

}

func strStr(haystack string, needle string) int {
	lh := len(haystack)
	ln := len(needle)
	if lh < ln {
		return -1
	}

	for pj := 0; pj+ln <= lh; pj++ {

		if haystack[pj:pj+ln] == needle {
			return pj
		}

	}
	return -1

}

func countAndSay(n int) string {
	if n == 1 {
		return "1"
	} else if n == 2 {
		return "11"
	} else if n == 3 {
		return "21"
	} else if n == 4 {
		return "1211"
	} else {
		s1 := countAndSay(n - 1)
		// 定义结果
		result := ""
		// 对s1遍历处理获取值
		local := s1[0:1]
		count := 0
		for i := 0; i < len(s1); i++ {
			// 设定计数器 计算同一个数字出现的次数 count
			if s1[i:i+1] == local {
				count++
			} else {

				result = result + fmt.Sprintf("%d", count) + local
				count = 1
				local = s1[i : i+1]
			}
		}
		result = result + fmt.Sprintf("%d", count) + local
		return result

	}

}

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	arrLen := len(strs)
	aj := 0
	ajl := len(strs[0])

	for i := 1; i < arrLen; i++ {
		if len(strs[i]) < ajl {
			aj = i
			ajl = len(strs[i])
		}
	}

	vmx := len(strs[aj])
	fmt.Printf("Lx :%d\n", aj)

	for ps := vmx; ps >= 0; ps-- {

		cpre := strs[aj][0:ps]
		lx := len(cpre)
		counter := 0
		fmt.Printf("MaxSub:%d\n", lx)

		for pj := 0; pj < arrLen; pj++ {
			if pj == aj {
				counter++
				continue
			}

			if strs[pj][0:lx] == cpre {
				counter++
			} else {
				break
			}

		}

		if counter == arrLen && counter > 1 {
			return cpre
		}

	}

	return ""

}

func removeNthFromEnd(head *ListNode, n int) *ListNode {

	pFirst := head
	pSecond := head

	if head.Next == nil && n >= 1 {

		return nil
	}
	//-- Forward N steps
	for i := 0; i < n; i++ {
		pFirst = pFirst.Next
	}

	if pFirst == nil {
		head = head.Next
	} else {

		for pFirst.Next != nil {

			pFirst = pFirst.Next
			pSecond = pSecond.Next
		}
	}
	pSecond.Next = pSecond.Next.Next

	return head

}
func reverseList(head *ListNode) *ListNode {

	pH := head
	pS := head.Next
	pH.Next = nil
	np := pH

	for pS != nil {
		np = pS
		pS = pS.Next
		np.Next = pH
		pH = np

	}
	return pH

}
func pri(head *ListNode) {

	for head != nil {

		fmt.Printf("%d,", head.Val)
		head = head.Next
	}
	fmt.Println()

}

func isPalindromeArr(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	fast, slow := head, head

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
	}
	if fast != nil {
		//---奇数个链表
		slow = slow.Next
	}

	s1 := 0
	s2 := 0
	t := 1

	for {

		if slow == nil {

			break
		}

		s1 = s1*10 + slow.Val

		s2 = s2 + t*head.Val

		t = t * 10

		slow = slow.Next
		head = head.Next

	}

	return s1 == s2

}
func hasCycle(head *ListNode) bool {

	if head == nil || head.Next == nil {
		return false
	}

	fast, slow := head, head

	//  fast 走2 步
	//  slow 走1 不
	//   只要有环,一定会遇到

	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			return true
		}
	}

	return false

}

func LevelOrder(root *TreeNode) [][]int {

	///------------没有分层: 错误的发给你是
	if root == nil {
		return nil
	}

	var result [][]int = nil
	var lQueue []*TreeNode = nil

	lQueue = append(lQueue, root)

	for {
		if len(lQueue) == 0 {
			break
		}

		cNode := lQueue[0]

		var cr []int = nil
		cr = append(cr, cNode.Val)

		if cNode.Left != nil {
			lQueue = append(lQueue, cNode.Left)
		}
		if cNode.Right != nil {
			lQueue = append(lQueue, cNode.Right)
		}
		lQueue = lQueue[1:]

		result = append(result, cr)

	}

	return result

}
func LevelOrderX(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}

	var queue []*TreeNode = nil
	queue = append(queue, root)
	//注意:----层分隔符
	queue = append(queue, nil)

	var res [][]int = nil

	var cp []int = nil
	//var level int = 1
	for _, tn := range queue {
		if tn != nil {
			//fmt.Printf("%d:[%d,%d],", level, index, tn.Val)
			cp = append(cp, tn.Val)
		}
	}
	res = append(res, cp)

	for len(queue) > 0 {
		var cNode = queue[0]
		queue = queue[1:]
		if cNode == nil {
			if len(queue) > 0 {
				//----层分隔
				//level = level + 1
				queue = append(queue, nil)
				var row []int = nil
				for _, tn := range queue {
					if tn != nil {

						row = append(row, tn.Val)
					}
				}
				res = append(res, row)
			}
		} else {

			if cNode.Left != nil {
				queue = append(queue, cNode.Left)
			}
			if cNode.Right != nil {
				queue = append(queue, cNode.Right)
			}
		}

	}

	return res

}

func TestLevelOrder() {
	//3,9,20,null,null,15,7

	var a15 = &TreeNode{Val: 15}
	var a7 = &TreeNode{Val: 7}

	var a18 = &TreeNode{Val: 18}
	var a27 = &TreeNode{Val: 27}

	var a20 = &TreeNode{Val: 20, Left: a15, Right: a7}

	var a9 = &TreeNode{Val: 9, Left: a18, Right: a27}

	var a3 = &TreeNode{Val: 3, Left: a9, Right: a20}

	fmt.Printf("%v\n", a3)
	var lo = LevelOrder(a3)

	fmt.Printf("\n%v\n", lo)

	lo = LevelOrderX(a3)

	fmt.Printf("\n%v\n", lo)

}
func testArray() {
	var nums = []int{1, 2, 3, 4}
	var t = SortedArrayToBST(nums)

	fmt.Printf("%v", t)
}

func testMegerNums() {

	// var nums1 = []int{0}
	// var nusm2 = []int{1}
	// Merge(nums1, 0, nusm2, 1)

	// fmt.Printf("%v\n", nums1)

	// var nums11 = []int{1}
	// var nusm21 = []int{0}
	// Merge(nums11, 1, nusm21, 0)

	// fmt.Printf("%v\n", nums11)

	// var nums31 = []int{1, 2, 3, 0, 0, 0}
	// var nusm32 = []int{2, 5, 6}
	// Merge(nums31, 3, nusm32, 3)

	// fmt.Printf("%v\n", nums31)

	var nums41 = []int{2, 0}
	var nusm42 = []int{1}

	fmt.Printf("%v\n", nums41)
	fmt.Printf("%v\n", nusm42)
	Merge(nums41, 1, nusm42, 1)

	fmt.Printf("%v\n", nums41)

}

func testMaxSub() {

	var arr = []int{-2, 1, -3, 4, -1, 2, 1, -5, 4}

	var mg = MaxSubArray(arr)

	fmt.Printf("%d\n", mg)
}

func testRobMax() {

	var arr = []int{2, 7, 9, 3, 1}

	var mg = Rob(arr)

	fmt.Printf("%d\n", mg)
}

func countPrimes(n int) int {
	var counter = 0
	for k := 2; k <= n; k++ {
		if IsPrime(k) {
			counter++
		}
	}
	return counter

}

func fizzBuzz(n int) []string {
	var str = make([]string, n)

	for i := 1; i <= n; i++ {

		if i%3 == 0 && i%5 == 0 {
			str[i-1] = "Fizz"
		} else if i%3 == 0 {
			str[i-1] = "Buzz"
		} else if i%5 == 0 {
			str[i-1] = "FizzBuzz"
		} else {
			str[i-1] = fmt.Sprintf("%d", i)

		}

	}
	return str
}

func isPowerOfThree(n int) bool {
	if n == 1 {
		return true
	}

	for n%3 == 0 {
		n = n / 3
		if n == 1 || n == -1 {
			return true
		}
	}

	return false

}

var pop8tab = "" +
	"\x00\x01\x01\x02\x01\x02\x02\x03\x01\x02\x02\x03\x02\x03\x03\x04" +
	"\x01\x02\x02\x03\x02\x03\x03\x04\x02\x03\x03\x04\x03\x04\x04\x05" +
	"\x01\x02\x02\x03\x02\x03\x03\x04\x02\x03\x03\x04\x03\x04\x04\x05" +
	"\x02\x03\x03\x04\x03\x04\x04\x05\x03\x04\x04\x05\x04\x05\x05\x06" +
	"\x01\x02\x02\x03\x02\x03\x03\x04\x02\x03\x03\x04\x03\x04\x04\x05" +
	"\x02\x03\x03\x04\x03\x04\x04\x05\x03\x04\x04\x05\x04\x05\x05\x06" +
	"\x02\x03\x03\x04\x03\x04\x04\x05\x03\x04\x04\x05\x04\x05\x05\x06" +
	"\x03\x04\x04\x05\x04\x05\x05\x06\x04\x05\x05\x06\x05\x06\x06\x07" +
	"\x01\x02\x02\x03\x02\x03\x03\x04\x02\x03\x03\x04\x03\x04\x04\x05" +
	"\x02\x03\x03\x04\x03\x04\x04\x05\x03\x04\x04\x05\x04\x05\x05\x06" +
	"\x02\x03\x03\x04\x03\x04\x04\x05\x03\x04\x04\x05\x04\x05\x05\x06" +
	"\x03\x04\x04\x05\x04\x05\x05\x06\x04\x05\x05\x06\x05\x06\x06\x07" +
	"\x02\x03\x03\x04\x03\x04\x04\x05\x03\x04\x04\x05\x04\x05\x05\x06" +
	"\x03\x04\x04\x05\x04\x05\x05\x06\x04\x05\x05\x06\x05\x06\x06\x07" +
	"\x03\x04\x04\x05\x04\x05\x05\x06\x04\x05\x05\x06\x05\x06\x06\x07" +
	"\x04\x05\x05\x06\x05\x06\x06\x07\x05\x06\x06\x07\x06\x07\x07\x08"

func generate(numRows int) [][]int {
	var res [][]int = nil

	for i := 0; i < numRows; i++ {
		var crow = make([]int, i+1)

		for j := 0; j < i+1; j++ {

			if i >= 2 && j >= 1 && j < i {
				crow[j] = res[i-1][j-1] + res[i-1][j]
			} else {
				crow[j] = 1
			}

		}

		// if i >= 2 {

		// 	var preRow = res[i-1]

		// 	for j := 1; j < i; j++ {

		// 		crow[j] = preRow[j-1] + preRow[j]
		// 	}
		// }

		res = append(res, crow)
	}

	return res

}
func isValid(s string) bool {
	var lx = len(s)

	if lx == 0 || lx%2 == 1 {
		//空字符或是奇数长度肯定不合法
		return false
	}

	var stack []byte = make([]byte, lx/2+1)
	stack[0] = s[0]
	var pStack = 0
	for i := 1; i < lx; i++ {

		if pStack == lx/2 {
			return false
		}

		var cb = s[i]
		if pStack < 0 {

			stack[0] = cb
			pStack = 0
			continue
		}
		var pb = stack[pStack]
		if (pb == '(' && cb == ')') || (pb == '[' && cb == ']') || (pb == '{' && cb == '}') {
			//---配对的
			pStack = pStack - 1
		} else {
			stack[pStack+1] = cb
			pStack = pStack + 1
		}
	}

	return pStack == -1

}

func missingNumber(nums []int) int {

	var lx = len(nums)
	var sum1 int = 0
	var sum2 int = 0
	//---原理说明：
	// [3,0,1]
	//----[0..n]中的 N    个数如果全部都在的话就是： 0 + ...+N

	for i := 0; i < lx; i++ {

		sum1 = sum1 + nums[i]
		sum2 = sum2 + (i + 1)

	}

	return sum2 - sum1

}

func threeSum(nums []int) [][]int {

	var res [][]int = nil
	var lx = len(nums)

	if lx < 3 {
		return nil
	}
	if lx == 3 {
		if (nums[0] + nums[1] + nums[2]) == 0 {
			return [][]int{
				nums,
			}
		} else {
			return nil
		}
	}

	sort.Ints(nums)
	fmt.Printf("%v\n", nums)
	for i := 0; i < lx-2; i++ {
		var fixNum = nums[i]

		if i >= 1 && fixNum == nums[i-1] {
			continue
		}

		if fixNum > 0 {
			//----因为已经拍好序了从小到大
			break
		}
		var ls = 0 - fixNum
		//--- 只有找到    pL + pR  == ls  就行了
		var pL = i + 1
		var pR = lx - 1

		for pL < pR && pR < lx {

			var cLR = nums[pL] + nums[pR]
			// fmt.Printf("pL=%d:%d,pR=%d,%d \n ", pL, nums[pL], pR, nums[pR])
			if cLR == ls {
				//---找到了符合要求的值
				if len(res) >= 1 {
					var lr = res[len(res)-1]
					if lr[0] == fixNum && lr[1] == nums[pL] && lr[2] == nums[pR] {
						//---not append
					} else {
						res = append(res, []int{fixNum, nums[pL], nums[pR]})
					}
				} else {
					res = append(res, []int{fixNum, nums[pL], nums[pR]})
				}

				pL = pL + 1
				pR = pR - 1

			} else if cLR > ls {
				pR = pR - 1
				//-----右边的是大书

			} else {
				pL = pL + 1

			}
		}
	}
	return res
}

func GroupAnagrams(strs []string) [][]string {
	// 	给你一个字符串数组，请你将 字母异位词 组合在一起。可以按任意顺序返回结果列表。
	// 字母异位词 是由重新排列源单词的字母得到的一个新单词，所有源单词中的字母都恰好只用一次。
	//
	// 示例 1:
	// 输入: strs = ["eat", "tea", "tan", "ate", "nat", "bat"]
	// 输出: [["bat"],["nat","tan"],["ate","eat","tea"]]

	strlx := len(strs)
	if strlx == 0 {
		return nil
	}
	if strlx == 0 {

		return [][]string{
			{strs[0]},
		}
	}

	var res [][]string = nil
	var ks = make(map[string][]string, strlx)

	for i := 0; i < strlx; i++ {

		var key = strs[i]
		var data = []rune(key)
		sort.Slice(data, func(i, j int) bool {
			return data[i] < data[j]
		})
		var sk = string(data)
		ks[sk] = append(ks[sk], strs[i])
	}
	for k, v := range ks {

		fmt.Printf("%v\n", k)
		res = append(res, v)
	}
	fmt.Printf("%v\n", res)

	return res

}
