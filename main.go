package main

import (
	"fmt"
	"strconv"
)

func main() {

	// 自定义排序
	//gramma_learning.SortSelfDefine()

	// 字符串和字节的转换
	//gramma_learning.ByteAndRune()

	if atoi, err := strconv.Atoi("+"); err == nil {
		fmt.Println(atoi)

	}

}
