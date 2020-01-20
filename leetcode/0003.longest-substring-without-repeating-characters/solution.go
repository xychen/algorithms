package main

import (
	"fmt"
)

// leetCode官方解答： https://leetcode-cn.com/articles/longest-substring-without-repeating-characters/

// 滑动窗口
func lengthOfLongestSubstring(s string) int {
	ht := make(map[byte]int)
	ans := 0
	low := 0
	high := 0
	n := len(s)
	for low < n && high < n {
		char := s[high]
		if _, ok := ht[char]; ok { //出现重复
			lowChar := s[low]
			delete(ht, lowChar)
			low += 1
		} else { //未出现重复
			ht[char] = high
			high += 1
			if high-low > ans {
				ans = high - low
			}
		}
	}
	return ans
}

//优化的滑动窗口
//当有重复的时候，左边界不需要一次只移动一个，可以直接跳到重复元素的下一个
func lengthOfLongestSubstring2(s string) int {
	ht := make(map[byte]int)
	ans := 0
	low := 0
	high := 0
	n := len(s)
	for high < n {
		char := s[high]
		if i, ok := ht[char]; ok && i >= low { //存在重复的字符串
			low = i + 1
		} else { //不存在重复的字符串
			ht[char] = high
			high += 1
			if high-low > ans {
				ans = high - low
			}
		}
	}
	return ans
}

func main() {
	str := "pwwkew"
	res := lengthOfLongestSubstring(str)
	res2 := lengthOfLongestSubstring2(str)
	fmt.Println(res, res2)
}
