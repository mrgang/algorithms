package main

import (
	"errors"
	"fmt"
	"math"
	"strings"
)

func main() {
	fmt.Println("algorithm ...")

	fmt.Println(longestPalindrome("cbbd"))

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

	return s
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
