package main

import "fmt"

/**
* @Author: super
* @Date: 2021-03-19 08:07
* @Description:
    字符串排序：输入：字符串由数字、小写字母、大写字母组成。输出：排序好的字符串。
    排序的标准：1. 数字>小写字母>大写字母。 2. 数字、字母间的相对顺序不变。 3. 额外存储空间：O（1）
    // Example
    input: "abcd4312ABDC"
    output: "4312abcdABDC"
**/

func sortString(str string) string {
	arr := make([]uint8, len(str))
	for i := 0; i < len(str); i++ {
		arr[i] = str[i]
	}
	index, curr := 0, 0
	for curr < len(str) {
		if arr[curr] >= '0' && arr[curr] <= '9' {
			adjust(index, curr, arr)
			index++
		}
		curr++
	}
	curr = index
	for curr < len(str) {
		if arr[curr] >= 'a' && arr[curr] <= 'z' {
			adjust(index, curr, arr)
			index++
		}
		curr++
	}
	return string(arr)
}

func adjust(index int, curr int, arr []uint8) {
	temp := arr[curr]
	for i := curr; i > index; i-- {
		arr[i] = arr[i-1]
	}
	arr[index] = temp
}

func main() {
	fmt.Println(sortString("1ASF23aswqedqdasfADS213AD1234AS"))
}