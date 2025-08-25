package main

import (
	"fmt"
)

// 只出现一次的数字：给定一个非空整数数组，除了某个元素只出现一次以外，其余每个元素均出现两次。找出那个只出现了一次的元素。可以使用 for 循环遍历数组，结合 if 条件判断和 map 数据结构来解决，例如通过 map 记录每个元素出现的次数，然后再遍历 map 找到出现次数为1的元素。
func SingleNumber(nums []int) int {

	var m = make(map[int]int)
	for _, num := range nums {
		m[num]++
	}
	for k, v := range m {
		if v == 1 {
			return k
		}
	}
	return -1
}

func main() {
	nums := []int{4, 1, 2, 1, 2}
	result := SingleNumber(nums)
	fmt.Println(result) // 输出 4
}

// isVaild 判断括号字符串是否有效
func isValid(s string) bool {

	//创建一个映射表 用于匹配右括号对应的左括号
	mapping := map[rune]rune{
		')': '(',
		']': '[',
		'}': '{',
	}

	//创建一个栈 用于存储左括号
	stack := []rune{}

	// 遍历字符串中的每个字符
	for _, chart := range s {
		// 如果字符是右括号,
		if _, ok := mapping[chart]; ok {
			// 如果是右括号，但栈为空，说明没有对应的左括号
			if len(stack) == 0 {
				return false
			}

			// 弹出栈顶元素， 检查是否匹配

		} else {
			// 如果字符是左括号，将其压入栈中
			stack = append(stack, chart)
		}
	}

}
