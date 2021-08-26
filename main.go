package main

import (
	"bytes"
	"errors"
	"fmt"
	"math"
	"strings"
)

func main() {

	fmt.Println(CRC16_Modbus([]uint8{
		1, 1, 1, 1, 1,
	}, 5), 0xAC89)
}

/*
1. 两数之和
给定一个整数数组 nums 和一个整数目标值 target，请你在该数组中找出和为目标值的那两个整数，并返回它们的数组下标。
你可以假设每种输入只会对应一个答案。但是，数组中同一个元素在答案里不能重复出现。
你可以按任意顺序返回答案。
输入：nums = [2,7,11,15], target = 9
输出：[0,1]
解释：因为 nums[0] + nums[1] == 9 ，返回 [0, 1] 。
fmt.Printf("%+v",twoSum([]int{2,7,11,15},9))
*/
func twoSum(nums []int, target int) []int {
	var mapv = make(map[int]interface{})
	for i, v := range nums {
		needValue := target - v
		if mapv[needValue] == nil {
			mapv[v] = i
		} else {
			return []int{i, mapv[needValue].(int)}
		}

	}
	return nil
}

/*
2. 两数相加
给你两个非空 的链表，表示两个非负的整数。它们每位数字都是按照逆序的方式存储的，并且每个节点只能存储一位数字。
请你将两个数相加，并以相同形式返回一个表示和的链表。
你可以假设除了数字 0 之外，这两个数都不会以 0开头。
输入：l1 = [2,4,3], l2 = [5,6,4]
输出：[7,0,8]
解释：342 + 465 = 807.
*/
type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var retList *ListNode
	var headList *ListNode
	var addOne = 0
	for l1 != nil || l2 != nil {
		var tmpSum = addOne
		if l1 != nil {
			tmpSum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			tmpSum += l2.Val
			l2 = l2.Next
		}
		addOne = tmpSum / 10
		if retList != nil {
			retList.Next = &ListNode{tmpSum % 10, nil}
			retList = retList.Next
		} else {
			headList = &ListNode{tmpSum % 10, nil}
			retList = headList
		}
	}
	if addOne == 1 {
		retList.Next = &ListNode{1, nil}
	}
	return headList
}

/*
3. 无重复字符的最长子串
给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度
输入: s = "abcabcbb"
输出: 3
解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
*/
func lengthOfLongestSubstring(s string) int {
	runeString := []rune(s)
	maxCount := 0
	subString := ""
	for _, intItem := range runeString {
		charItem := fmt.Sprintf("%c", intItem)
		sidx := strings.Index(subString, charItem)
		if sidx == -1 {
			subString += charItem
		} else {
			maxCount = int(math.Max(float64(len(subString)), float64(maxCount)))
			subString = subString[sidx+1:]
			subString += charItem
		}
	}
	return int(math.Max(float64(len(subString)), float64(maxCount)))
}

/*
4. 寻找两个正序数组的中位数
给定两个大小分别为 m 和 n 的正序（从小到大）数组 nums1 和 nums2。请你找出并返回这两个正序数组的 中位数 。
输入：nums1 = [1,3], nums2 = [2]
输出：2.00000
解释：合并数组 = [1,2,3] ，中位数 2
fmt.Println(findMedianSortedArrays([]int{1,2},[]int{3,4}))
*/
func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	m := len(nums1)
	n := len(nums2)

	if (m + n) == 0 {
		return 0
	}
	if n == 0 {
		if m%2 == 1 {
			return float64(nums1[m/2])
		} else {
			return float64(nums1[m/2]+nums1[(m-2)/2]) / 2
		}
	}
	if m == 0 {
		if n%2 == 1 {
			return float64(nums2[n/2])
		} else {
			return float64(nums2[n/2]+nums2[(n-2)/2]) / 2
		}
	}
	return float64(getIndexValue(&nums1, &nums2, (m+n+1)/2)+getIndexValue(&nums1, &nums2, (m+n+2)/2)) / 2

}
func getIndexValue(nums1 *[]int, nums2 *[]int, centerIdx int) int {
	nums1Idx, nums2Idx := 0, 0
	hasAdd1, hasAdd2 := false, false
	centValue := 0
	for i := 0; i < centerIdx; i++ {
		err1, v1 := getArrayValue(nums1, nums1Idx)
		err2, v2 := getArrayValue(nums2, nums2Idx)
		if err1 != nil {
			return (*nums2)[nums2Idx+centerIdx-i-1]
		}
		if err2 != nil {
			return (*nums1)[nums1Idx+centerIdx-i-1]
		}
		if v1 > v2 {
			if hasAdd2 {
				hasAdd2 = true
			} else {
				nums2Idx++
			}
			centValue = v2
		} else {
			if hasAdd1 {
				hasAdd1 = true
			} else {
				nums1Idx++
			}
			centValue = v1
		}
	}
	return centValue

}
func getArrayValue(nums *[]int, idx int) (error, int) {
	if idx < len(*nums) {
		return nil, (*nums)[idx]
	} else {
		return errors.New("no value"), 0
	}
}

/*
5. 最长回文子串
给你一个字符串 s，找到 s 中最长的回文子串。
*/
func longestPalindrome(s string) string {
	len := len(s)
	if len < 2 {
		return s
	}
	left, right := 0, 0

	for i, _ := range s {
		//以i为中心
		left1, right1 := i, i
		hasChange := false
		for left1 >= 0 && (right1 <= len-1) && s[left1] == s[right1] {
			left1--
			right1++
			hasChange = true
		}
		if hasChange {
			left1++
			right1--
		}
		if right1-left1 > right-left {
			left = left1
			right = right1
		}

		if i >= 0 && i+1 <= len-1 {
			//以i,i+1为中心
			left2, right2 := i, i+1
			hasChange := false
			for left2 >= 0 && (right2 <= len-1) && s[left2] == s[right2] {
				left2--
				right2++
				hasChange = true
			}
			if hasChange {
				left2++
				right2--
			} else {
				left2, right2 = i, i
			}
			if right2-left2 > right-left {
				left = left2
				right = right2
			}
		}

	}
	return s[left : right+1]
}

/*
6. Z 字形变换
将一个给定字符串 s 根据给定的行数 numRows ，以从上往下、从左到右进行 Z 字形排列。
比如输入字符串为 "PAYPALISHIRING" 行数为 3 时，排列如下：
P   A   H   N
A P L S I I G
Y   I   R
之后，你的输出需要从左往右逐行读取，产生出一个新的字符串，比如："PAHNAPLSIIGYIR"。
*/
func convert(s string, numRows int) string {
	len := len(s)
	if len < 2 {
		return s
	}
	numbers := [][]string{}
	marked := 0
	numCols := 0
	for {
		//向下
		var numbersCol []string
		added := 0
		for i := 0; i < numRows; i++ {
			numbersCol = append(numbersCol, string(s[marked]))
			marked++
			added++
			if marked == len {
				break
			}
		}

		empl := numRows - added
		for i := 0; i < empl; i++ {
			numbersCol = append(numbersCol, "")
		}
		numbers = append(numbers, numbersCol)
		numCols++
		if marked == len {
			break
		}
		//向上
		var numbersCol2 []string
		added = 0
		for i := numRows - 2; i > 0; i-- {
			numbersCol2 = append(numbersCol2, string(s[marked]))
			marked++
			added++
			if marked == len {
				break
			}
		}
		empl = numRows - added
		var emparr []string
		for i := 0; i < empl-1; i++ {
			emparr = append(emparr, "")
		}
		for i := added - 1; i >= 0; i-- {
			emparr = append(emparr, numbersCol2[i])
		}
		numbersCol2 = append(emparr, "")
		numbers = append(numbers, numbersCol2)
		numCols++
		if marked == len {
			break
		}
	}
	var buffer bytes.Buffer
	for i := 0; i < numRows; i++ {
		for j := 0; j < numCols; j++ {
			buffer.WriteString(numbers[j][i])
		}
	}
	return buffer.String()
}

/*7. 整数反转
给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
假设环境不允许存储 64 位整数（有符号或无符号）。
示例 1：
输入：x = 123
输出：321
*/
func reverse(x int) int {
	var isPositive = true
	if x > 0 {
		isPositive = true
	} else {
		isPositive = false
		x = -x
	}
	retInt := 0
	for x != 0 {
		lastInt := x % 10
		x = x / 10
		retInt = retInt*10 + lastInt
	}
	if retInt > 2147483647 {
		return 0
	}
	if isPositive {
		return retInt
	} else {
		return -retInt
	}
}

/*8. 字符串转换整数 (atoi)
请你来实现一个myAtoi(string s)函数，使其能将字符串转换成一个 32 位有符号整数（类似 C/C++ 中的 atoi 函数）。

函数myAtoi(string s) 的算法如下：

读入字符串并丢弃无用的前导空格
检查下一个字符（假设还未到字符末尾）为正还是负号，读取该字符（如果有）。 确定最终结果是负数还是正数。 如果两者都不存在，则假定结果为正。
读入下一个字符，直到到达下一个非数字字符或到达输入的结尾。字符串的其余部分将被忽略。
将前面步骤读入的这些数字转换为整数（即，"123" -> 123， "0032" -> 32）。如果没有读入数字，则整数为 0 。必要时更改符号（从步骤 2 开始）。
如果整数数超过 32 位有符号整数范围 [−231, 231− 1] ，需要截断这个整数，使其保持在这个范围内。具体来说，小于 −231 的整数应该被固定为 −231 ，大于 231 − 1 的整数应该被固定为 231 − 1 。
返回整数作为最终结果。
注意：

本题中的空白字符只包括空格字符 ' ' 。
除前导空格或数字后的其余字符串外，请勿忽略 任何其他字符。
*/
func myAtoi(s string) int {
	const maxInt int64 = 1 << 31
	var (
		i    int
		r    int64
		sign int64 = 1
	)
	for ; i < len(s); i++ {
		if s[i] != ' ' {
			break
		}
	}
	findSign := false
	if s[i] == '-' {
		sign = -1
		i++
		findSign = true
	}
	if s[i] == '+' {
		if findSign {
			return 0
		}
		sign = 1
		i++

	}

	for ; i < len(s) && s[i] >= '0' && s[i] <= '9'; i++ {
		r = r*10 + int64(s[i]-'0')
		if r >= maxInt {
			if sign == 1 {
				r = maxInt - 1
			} else {
				r = maxInt
			}
			break
		}
	}
	return int(r * sign)
}

/*9. 回文数
给你一个整数 x ，如果 x 是一个回文整数，返回 true ；否则，返回 false 。
回文数是指正序（从左向右）和倒序（从右向左）读都是一样的整数。例如，121 是回文，而 123 不是。
*/
func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}
	if x < 10 {
		return true
	}

	var (
		curInt     int = x
		lastInt    int
		reverseInt int
	)
	for curInt != 0 {
		lastInt = curInt % 10
		reverseInt = reverseInt*10 + lastInt
		curInt = curInt / 10
	}
	return reverseInt == x
}

/*10. 正则表达式匹配
给你一个字符串s和一个字符规律p，请你来实现一个支持 '.'和'*'的正则表达式匹配。
'.' 匹配任意单个字符
'*' 匹配零个或多个前面的那一个元素
所谓匹配，是要涵盖整个字符串s的，而不是部分字符串。
*/
func isMatch(s string, p string) bool {
	if len(p) > len(s) || len(p)*len(s) == 0 {
		return false
	}

	return false
}

/*709. 转换成小写字母
实现函数 ToLowerCase()，该函数接收一个字符串参数 str，并将该字符串中的大写字母转换成小写字母，之后返回新的字符串。
*/
func toLowerCase(str string) string {
	var result = strings.Builder{}
	for _, char := range str {
		if char >= 'A' && char <= 'Z' {
			result.WriteString(string(char - 'A' + 'a'))
		} else {
			result.WriteString(string(char))
		}

	}
	return result.String()
}

/* 快排
var testArr = []int{5,6,11,55,99,5,2,6,9,7,11,2,3,1,55}
	quickSort(&testArr,0,len(testArr)-1)
	fmt.Println(testArr)
*/
func quickSort(arr *[]int, left int, right int) {
	if left < right {
		var i = partition(arr, left, right)
		quickSort(arr, left, i-1)
		quickSort(arr, i+1, right)
	}
}
func partition(arr *[]int, left int, right int) int {
	var base = (*arr)[left]
	var i = left
	var j = right
	for i < j {
		for j > i && (*arr)[j] >= base {
			j--
		}
		for i < j && (*arr)[i] <= base {
			i++
		}
		(*arr)[i], (*arr)[j] = (*arr)[j], (*arr)[i]
	}
	(*arr)[i], (*arr)[left] = (*arr)[left], (*arr)[i]
	return i
}

// CRC 高位字节值表
var s_CRCHi = []byte{
	0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0,
	0x80, 0x41, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41,
	0x00, 0xC1, 0x81, 0x40, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0,
	0x80, 0x41, 0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1, 0x81, 0x40,
	0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1,
	0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0, 0x80, 0x41,
	0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1,
	0x81, 0x40, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41,
	0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0,
	0x80, 0x41, 0x00, 0xC1, 0x81, 0x40, 0x00, 0xC1, 0x81, 0x40,
	0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1,
	0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1, 0x81, 0x40,
	0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0,
	0x80, 0x41, 0x00, 0xC1, 0x81, 0x40, 0x00, 0xC1, 0x81, 0x40,
	0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0,
	0x80, 0x41, 0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1, 0x81, 0x40,
	0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0,
	0x80, 0x41, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41,
	0x00, 0xC1, 0x81, 0x40, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0,
	0x80, 0x41, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41,
	0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0,
	0x80, 0x41, 0x00, 0xC1, 0x81, 0x40, 0x00, 0xC1, 0x81, 0x40,
	0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0, 0x80, 0x41, 0x00, 0xC1,
	0x81, 0x40, 0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41,
	0x00, 0xC1, 0x81, 0x40, 0x01, 0xC0, 0x80, 0x41, 0x01, 0xC0,
	0x80, 0x41, 0x00, 0xC1, 0x81, 0x40}

// CRC 低位字节值表
var s_CRCLo = []byte{
	0x00, 0xC0, 0xC1, 0x01, 0xC3, 0x03, 0x02, 0xC2, 0xC6, 0x06,
	0x07, 0xC7, 0x05, 0xC5, 0xC4, 0x04, 0xCC, 0x0C, 0x0D, 0xCD,
	0x0F, 0xCF, 0xCE, 0x0E, 0x0A, 0xCA, 0xCB, 0x0B, 0xC9, 0x09,
	0x08, 0xC8, 0xD8, 0x18, 0x19, 0xD9, 0x1B, 0xDB, 0xDA, 0x1A,
	0x1E, 0xDE, 0xDF, 0x1F, 0xDD, 0x1D, 0x1C, 0xDC, 0x14, 0xD4,
	0xD5, 0x15, 0xD7, 0x17, 0x16, 0xD6, 0xD2, 0x12, 0x13, 0xD3,
	0x11, 0xD1, 0xD0, 0x10, 0xF0, 0x30, 0x31, 0xF1, 0x33, 0xF3,
	0xF2, 0x32, 0x36, 0xF6, 0xF7, 0x37, 0xF5, 0x35, 0x34, 0xF4,
	0x3C, 0xFC, 0xFD, 0x3D, 0xFF, 0x3F, 0x3E, 0xFE, 0xFA, 0x3A,
	0x3B, 0xFB, 0x39, 0xF9, 0xF8, 0x38, 0x28, 0xE8, 0xE9, 0x29,
	0xEB, 0x2B, 0x2A, 0xEA, 0xEE, 0x2E, 0x2F, 0xEF, 0x2D, 0xED,
	0xEC, 0x2C, 0xE4, 0x24, 0x25, 0xE5, 0x27, 0xE7, 0xE6, 0x26,
	0x22, 0xE2, 0xE3, 0x23, 0xE1, 0x21, 0x20, 0xE0, 0xA0, 0x60,
	0x61, 0xA1, 0x63, 0xA3, 0xA2, 0x62, 0x66, 0xA6, 0xA7, 0x67,
	0xA5, 0x65, 0x64, 0xA4, 0x6C, 0xAC, 0xAD, 0x6D, 0xAF, 0x6F,
	0x6E, 0xAE, 0xAA, 0x6A, 0x6B, 0xAB, 0x69, 0xA9, 0xA8, 0x68,
	0x78, 0xB8, 0xB9, 0x79, 0xBB, 0x7B, 0x7A, 0xBA, 0xBE, 0x7E,
	0x7F, 0xBF, 0x7D, 0xBD, 0xBC, 0x7C, 0xB4, 0x74, 0x75, 0xB5,
	0x77, 0xB7, 0xB6, 0x76, 0x72, 0xB2, 0xB3, 0x73, 0xB1, 0x71,
	0x70, 0xB0, 0x50, 0x90, 0x91, 0x51, 0x93, 0x53, 0x52, 0x92,
	0x96, 0x56, 0x57, 0x97, 0x55, 0x95, 0x94, 0x54, 0x9C, 0x5C,
	0x5D, 0x9D, 0x5F, 0x9F, 0x9E, 0x5E, 0x5A, 0x9A, 0x9B, 0x5B,
	0x99, 0x59, 0x58, 0x98, 0x88, 0x48, 0x49, 0x89, 0x4B, 0x8B,
	0x8A, 0x4A, 0x4E, 0x8E, 0x8F, 0x4F, 0x8D, 0x4D, 0x4C, 0x8C,
	0x44, 0x84, 0x85, 0x45, 0x87, 0x47, 0x46, 0x86, 0x82, 0x42,
	0x43, 0x83, 0x41, 0x81, 0x80, 0x40}

func CRC16_Modbus(pBuf []uint8, uslen int) int {
	var ucCRCHi = 0xFF
	var ucCRCLo = 0xFF
	var usIndex = 0
	for i := 0; i < uslen; i++ {
		//usIndex = ucCRCHi ^ pBuf[i]
		//ucCRCHi = ucCRCLo ^ s_CRCHi[usIndex]
		//ucCRCLo = s_CRCLo[usIndex]
		usIndex = ucCRCLo ^ int(pBuf[i])
		ucCRCLo = ucCRCHi ^ int(s_CRCHi[usIndex])
		ucCRCHi = int(s_CRCLo[usIndex])
	}
	return ucCRCHi<<8 | ucCRCLo
}
