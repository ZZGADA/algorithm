package algorithm_go

import (
	"strings"
)

func convert(s string, numRows int) string {
	if numRows < 2 {
		return s
	}

	rows := make([][]string, numRows)
	for i := 0; i < numRows; i++ {
		rows[i] = make([]string, 0)
	}

	i := 0     // 模拟行
	flag := -1 // 模拟方向 和 步长
	for ci := range s {
		rows[i] = append(rows[i], string(s[ci]))
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}

	builder := new(strings.Builder)
	for j := 0; j < numRows; j++ {
		builder.WriteString(strings.Join(rows[j], ""))
	}
	return builder.String()

}

func quickConvert(s string, numRows int) string {
	if numRows < 2 || numRows >= len(s) {
		return s
	}

	// 使用strings.Builder代替[][]string，减少内存分配
	rows := make([]strings.Builder, numRows)
	i, flag := 0, -1

	// 直接操作字节，避免string转换
	for _, c := range s {
		rows[i].WriteRune(c)
		if i == 0 || i == numRows-1 {
			flag = -flag
		}
		i += flag
	}

	// 一次性分配足够大的缓冲区
	var result strings.Builder
	totalLen := 0
	for _, b := range rows {
		totalLen += b.Len()
	}
	result.Grow(totalLen) // 预分配内存，避免多次扩容

	// 拼接结果
	for _, b := range rows {
		result.WriteString(b.String())
	}

	return result.String()
}
